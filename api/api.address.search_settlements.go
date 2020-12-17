package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
)

const AddressMethodSearchSettlements = "searchSettlements"

type AddressSearchSettlements func(o ...func(*AddressSearchSettlementsRequest)) (*Response, error)

type AddressSearchSettlementsRequest struct {
	CityName string
	Limit    int

	ctx context.Context

	apiKey string
}

func newAddressSearchSettlementsFunc(t Transport, a string) AddressSearchSettlements {

	return func(o ...func(*AddressSearchSettlementsRequest)) (*Response, error) {
		var r = AddressSearchSettlementsRequest{apiKey: a}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// https://devcenter.novaposhta.ua/docs/services/556d7ccaa0fe4f08e8f7ce43/operations/58e5ebeceea27017bc851d67
func (r AddressSearchSettlementsRequest) Do(ctx context.Context, t Transport) (*Response, error) {

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
			"calledMethod": AddressMethodSearchSettlements,
			"methodProperties": map[string]interface{}{
				"CityName": r.CityName,
				"Limit":    r.Limit,
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

// WithContext sets the request context.
func (f AddressSearchSettlements) WithContext(v context.Context) func(*AddressSearchSettlementsRequest) {

	return func(r *AddressSearchSettlementsRequest) {
		r.ctx = v
	}
}

func (f AddressSearchSettlements) WithParams(cityName string, limit int) func(*AddressSearchSettlementsRequest) {

	return func(r *AddressSearchSettlementsRequest) {
		r.CityName = cityName
		r.Limit = limit
	}
}
