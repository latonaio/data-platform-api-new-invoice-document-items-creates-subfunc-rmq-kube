package requests

type AddressMaster struct {
	AddressID         int     `json:"AddressID"`
	ValidityEndDate   string  `json:"ValidityEndDate"`
	ValidityStartDate string  `json:"ValidityStartDate"`
	PostalCode        string  `json:"PostalCode"`
	LocalRegion       string  `json:"LocalRegion"`
	Country           string  `json:"Country"`
	GlobalRegion      string  `json:"GlobalRegion"`
	TimeZone          string  `json:"TimeZone"`
	District          *string `json:"District"`
	StreetName        string  `json:"StreetName"`
	CityName          string  `json:"CityName"`
	Building          *string `json:"Building"`
	Floor             *int    `json:"Floor"`
	Room              *int    `json:"Room"`
}
