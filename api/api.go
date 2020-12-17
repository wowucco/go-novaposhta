package api

import (
	"net/http"
)

type Transport interface {
	Perform(*http.Request) (*http.Response, error)
	IsJsonFormat() bool
	IsXmlFormat() bool
}

type API struct {
	Address *Address
}

func New(t Transport, apiKey string) *API {
	return &API{
		Address: &Address{
			GetCities:                newGetCitiesFunc(t, apiKey),
			GetWarehouses:            newGetWarehousesFunc(t, apiKey),
			SearchSettlements:        newAddressSearchSettlementsFunc(t, apiKey),
			SearchSettlementsStreets: newAddressSearchSettlementsStreetsFunc(t, apiKey),
		},
	}
}
