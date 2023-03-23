package requests

type CalculateAddressID struct {
	AddressIDLatestNumber *int `json:"AddressIDLatestNumber"`
	AddressID             int  `json:"AddressID"`
}
