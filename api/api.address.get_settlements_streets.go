package api

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"log"
)

const AddressMethodSearchSettlementsStreets = "searchSettlementStreets"

type AddressSearchSettlementsStreets func(o ...func(*AddressSearchSettlementsStreetsRequest)) (*Response, error)

type AddressSearchSettlementsStreetsRequest struct {
	StreetName    string
	SettlementRef string
	Limit         int

	ctx context.Context

	apiKey string
}

func newAddressSearchSettlementsStreetsFunc(t Transport, a string) AddressSearchSettlementsStreets {

	return func(o ...func(*AddressSearchSettlementsStreetsRequest)) (*Response, error) {
		var r = AddressSearchSettlementsStreetsRequest{apiKey: a}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// https://devcenter.novaposhta.ua/docs/services/556d7ccaa0fe4f08e8f7ce43/operations/58e5f369eea27017540b58ac
func (r AddressSearchSettlementsStreetsRequest) Do(ctx context.Context, t Transport) (*Response, error) {

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
			"calledMethod": AddressMethodSearchSettlementsStreets,
			"methodProperties": map[string]interface{}{
				"StreetName":    r.StreetName,
				"SettlementRef": r.SettlementRef,
				"Limit":         r.Limit,
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

func (f AddressSearchSettlementsStreets) WithContext(v context.Context) func(request *AddressSearchSettlementsStreetsRequest) {

	return func(r *AddressSearchSettlementsStreetsRequest) {
		r.ctx = v
	}
}

func (f AddressSearchSettlementsStreets) WithParams(streetName, settlementRef string, limit int) func(*AddressSearchSettlementsStreetsRequest) {

	return func(r *AddressSearchSettlementsStreetsRequest) {
		r.StreetName = streetName
		r.SettlementRef = settlementRef
		r.Limit = limit
	}
}
