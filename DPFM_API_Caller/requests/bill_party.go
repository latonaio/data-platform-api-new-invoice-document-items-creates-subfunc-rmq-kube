package requests

type BillParty struct {
	BillFromParty        int `json:"BillFromParty"`
	BillToParty          int `json:"BillToParty"`
	OrderID              int `json:"OrderID"`
	OrderItem            int `json:"OrderItem"`
	DeliveryDocument     int `json:"DeliveryDocument"`
	DeliveryDocumentItem int `json:"DeliveryDocumentItem"`
}
