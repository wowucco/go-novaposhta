package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
)

const AddressMethodGetWarehouses = "getWarehouses"

type AddressGetWarehouses func(o ...func(*AddressGetWarehousesRequest)) (*Response, error)

type AddressGetWarehousesRequest struct {
	CityName string // Дополнительный фильтр по имени города
	CityRef  string // Дополнительный фильтр по идентификатору города
	Page     int    // Страница, максимум 500 записей на странице. Работает в связке с параметром Limit
	Limit    int    // Количество записей на странице. Работает в связке с параметром Page
	Language string // Вывод описания на Украинском или русском языках - ru. По умолчанию всегда выводиться на Украинском языке.

	ctx context.Context

	apiKey string
}

func newGetWarehousesFunc(t Transport, a string) AddressGetWarehouses {

	return func(o ...func(*AddressGetWarehousesRequest)) (*Response, error) {
		var r = AddressGetWarehousesRequest{apiKey: a}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

func (r AddressGetWarehousesRequest) Do(ctx context.Context, t Transport) (*Response, error) {

	var (
		buf    bytes.Buffer
		method string
		params map[string]interface{}
	)

	method = "POST"

	if t.IsJsonFormat() {

		params = map[string]interface{}{
			"apiKey":       r.apiKey,
			"modelName":    AddressModelName,
			"calledMethod": AddressMethodGetWarehouses,
			"methodProperties": map[string]interface{}{
				"CityName": r.CityName,
				"CityRef":  r.CityRef,
				"Page":     r.Page,
				"Limit":    r.Limit,
				"Language": r.Language,
			},
		}

		if err := json.NewEncoder(&buf).Encode(params); err != nil {
			log.Fatalf("Error encoding json query: %s", err)
			return nil, err
		}
	} else {
		var params string

		if err := xml.NewEncoder(&buf).Encode(params); err != nil {
			log.Fatalf("Error encoding xml query: %s", err)
			return nil, err
		}
	}

	req, _ := newRequest(method, &buf)

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := t.Perform(req)

	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

func (r AddressGetWarehouses) WithContext(v context.Context) func(*AddressGetCitiesRequest) {

	return func(r *AddressGetCitiesRequest) {
		r.ctx = v
	}
}

func (r AddressGetWarehouses) WithParams(cityName, cityRef, language string, page, limit int) func(*AddressGetWarehousesRequest) {

	return func(r *AddressGetWarehousesRequest) {
		r.CityName = cityName
		r.CityRef = cityRef
		r.Language = language
		r.Page = page
		r.Limit = limit
	}
}
