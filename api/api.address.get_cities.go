package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
)

const AddressMethodGetCities = "getCities"

type AddressGetCities func(o ...func(*AddressGetCitiesRequest)) (*Response, error)

type AddressGetCitiesRequest struct {

	Ref string // Идентификатор города
	Page int   // Номер страницы для отображения
	FindByString string // Поиск по названию города

	ctx context.Context

	apiKey string
}

func newGetCitiesFunc(t Transport, a string) AddressGetCities {

	return func(o ...func(*AddressGetCitiesRequest)) (*Response, error) {
		var r = AddressGetCitiesRequest{apiKey: a}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

func (r AddressGetCitiesRequest) Do(ctx context.Context, t Transport) (*Response, error) {

	var (
		buf	   bytes.Buffer
		method string
		params map[string]interface{}
	)

	method = "POST"

	if t.IsJsonFormat() {

		params = map[string]interface{}{
			"apiKey":       r.apiKey,
			"modelName":    AddressModelName,
			"calledMethod": AddressMethodGetCities,
			"methodProperties": map[string]interface{}{
				"Ref":          r.Ref,
				"Page":         r.Page,
				"FindByString": r.FindByString,
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

func (r AddressGetCities) WithContext(v context.Context) func(*AddressGetCitiesRequest) {

	return func(r *AddressGetCitiesRequest) {
		r.ctx = v
	}
}

func (r AddressGetCities) WithParams(findByString, ref string, page int) func(*AddressGetCitiesRequest) {

	return func(r *AddressGetCitiesRequest) {
		r.Ref = ref
		r.Page = page
		r.FindByString = findByString
	}
}
