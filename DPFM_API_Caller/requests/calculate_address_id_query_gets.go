package requests

type CalculateAddressIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             *int   `json:"LatestNumber"`
}
