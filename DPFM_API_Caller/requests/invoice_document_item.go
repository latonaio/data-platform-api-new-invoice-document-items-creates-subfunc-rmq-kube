package requests

type InvoiceDocumentItem struct {
	OrderID                   int `json:"OrderID"`
	OrderItem                 int `json:"OrderItem"`
	DeliveryDocument          int `json:"DeliveryDocument"`
	DeliveryDocumentItem      int `json:"DeliveryDocumentItem"`
	InvoiceDocumentItemNumber int `json:"InvoiceDocumentItemNumber"`
}
