package api

const AddressModelName = "Address"

type Address struct {
	GetCities                AddressGetCities
	GetWarehouses            AddressGetWarehouses
	SearchSettlements        AddressSearchSettlements
	SearchSettlementsStreets AddressSearchSettlementsStreets
}
