package api_input_reader

type SDC struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	Filepath          string          `json:"filepath"`
	APIStatusCode     int             `json:"api_status_code"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	ServiceLabel      string          `json:"service_label"`
	InputParameters   InputParameters `json:"InputParameters"`
	Header            Header          `json:"InvoiceDocument"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	InvoiceDocumentID *int            `json:"invoice_document_id"`
	Deleted           bool            `json:"deleted"`
}

type InputParameters struct {
	InvoiceDocumentDate       *[]*string `json:"InvoiceDocumentDate"`
	InvoiceDocumentDateTo     *string    `json:"InvoiceDocumentDateTo"`
	InvoiceDocumentDateFrom   *string    `json:"InvoiceDocumentDateFrom"`
	BillFromParty             *[]*int    `json:"BillFromParty"`
	BillFromPartyTo           *int       `json:"BillFromPartyTo"`
	BillFromPartyFrom         *int       `json:"BillFromPartyFrom"`
	BillToParty               *[]*int    `json:"BillToParty"`
	BillToPartyTo             *int       `json:"BillToPartyTo"`
	BillToPartyFrom           *int       `json:"BillToPartyFrom"`
	ConfirmedDeliveryDate     *[]*string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateTo   *string    `json:"ConfirmedDeliveryDateTo"`
	ConfirmedDeliveryDateFrom *string    `json:"ConfirmedDeliveryDateFrom"`
	ActualGoodsIssueDate      *[]*string `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueDateTo    *string    `json:"ActualGoodsIssueDateTo"`
	ActualGoodsIssueDateFrom  *string    `json:"ActualGoodsIssueDateFrom"`
	ActualGoodsIssueTime      *[]string  `json:"ActualGoodsIssueTime"`
	ActualGoodsIssueTimeTo    string     `json:"ActualGoodsIssueTimeTo"`
	ActualGoodsIssueTimeFrom  string     `json:"ActualGoodsIssueTimeFrom"`
	ReferenceDocument         int        `json:"ReferenceDocument"`
}

type Header struct {
	InvoiceDocument                   int       `json:"InvoiceDocument"`
	CreationDate                      *string   `json:"CreationDate"`
	CreationTime                      *string   `json:"CreationTime"`
	LastChangeDate                    *string   `json:"LastChangeDate"`
	LastChangeTime                    *string   `json:"LastChangeTime"`
	SupplyChainRelationshipID         *int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID  *int      `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID  *int      `json:"SupplyChainRelationshipPaymentID"`
	BillToParty                       *int      `json:"BillToParty"`
	BillFromParty                     *int      `json:"BillFromParty"`
	BillToCountry                     *string   `json:"BillToCountry"`
	BillFromCountry                   *string   `json:"BillFromCountry"`
	Payer                             *int      `json:"Payer"`
	Payee                             *int      `json:"Payee"`
	InvoiceDocumentDate               *string   `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime               *string   `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate            *string   `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate              *string   `json:"InvoicePeriodEndDate"`
	AccountingPostingDate             *string   `json:"AccountingPostingDate"`
	IsExportImport                    *bool     `json:"IsExportImport"`
	HeaderBillingIsConfirmed          *bool     `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus           *string   `json:"HeaderBillingConfStatus"`
	TotalNetAmount                    *float32  `json:"TotalNetAmount"`
	TotalTaxAmount                    *float32  `json:"TotalTaxAmount"`
	TotalGrossAmount                  *float32  `json:"TotalGrossAmount"`
	TransactionCurrency               *string   `json:"TransactionCurrency"`
	Incoterms                         *string   `json:"Incoterms"`
	PaymentTerms                      *string   `json:"PaymentTerms"`
	DueCalculationBaseDate            *string   `json:"DueCalculationBaseDate"`
	PaymentDueDate                    *string   `json:"PaymentDueDate"`
	NetPaymentDays                    *int      `json:"NetPaymentDays"`
	PaymentMethod                     *string   `json:"PaymentMethod"`
	ExternalReferenceDocument         *string   `json:"ExternalReferenceDocument"`
	DocumentHeaderText                *string   `json:"DocumentHeaderText"`
	HeaderIsCleared                   *bool     `json:"HeaderIsCleared"`
	HeaderPaymentBlockStatus          *bool     `json:"HeaderPaymentBlockStatus"`
	HeaderPaymentRequisitionIsCreated *bool     `json:"HeaderPaymentRequisitionIsCreated"`
	IsCancelled                       *bool     `json:"IsCancelled"`
	Partner                           []Partner `json:"Partner"`
	Item                              []Item    `json:"InvoiceDocumentItem"`
	Address                           []Address `json:"Address"`
}

type Partner struct {
	InvoiceDocument         int     `json:"InvoiceDocument"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         int     `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type Address struct {
	InvoiceDocument int     `json:"InvoiceDocument"`
	AddressID       int     `json:"AddressID"`
	PostalCode      *string `json:"PostalCode"`
	LocalRegion     *string `json:"LocalRegion"`
	Country         *string `json:"Country"`
	District        *string `json:"District"`
	StreetName      *string `json:"StreetName"`
	CityName        *string `json:"CityName"`
	Building        *string `json:"Building"`
	Floor           *int    `json:"Floor"`
	Room            *int    `json:"Room"`
}

type Item struct {
	InvoiceDocument                         int                  `json:"InvoiceDocument"`
	InvoiceDocumentItem                     int                  `json:"InvoiceDocumentItem"`
	InvoiceDocumentItemCategory             *string              `json:"InvoiceDocumentItemCategory"`
	SupplyChainRelationshipID               *int                 `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       *int                 `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  *int                 `json:"SupplyChainRelationshipDeliveryPlantID"`
	InvoiceDocumentItemText                 *string              `json:"InvoiceDocumentItemText"`
	InvoiceDocumentItemTextByBuyer          *string              `json:"InvoiceDocumentItemTextByBuyer"`
	InvoiceDocumentItemTextBySeller         *string              `json:"InvoiceDocumentItemTextBySeller"`
	Product                                 *string              `json:"Product"`
	ProductGroup                            *string              `json:"ProductGroup"`
	ProductStandardID                       *string              `json:"ProductStandardID"`
	CreationDate                            *string              `json:"CreationDate"`
	CreationTime                            *string              `json:"CreationTime"`
	LastChangeDate                          *string              `json:"LastChangeDate"`
	LastChangeTime                          *string              `json:"LastChangeTime"`
	ItemBillingIsConfirmed                  *bool                `json:"ItemBillingIsConfirmed"`
	Buyer                                   *int                 `json:"Buyer"`
	Seller                                  *int                 `json:"Seller"`
	DeliverToParty                          *int                 `json:"DeliverToParty"`
	DeliverFromParty                        *int                 `json:"DeliverFromParty"`
	DeliverToPlant                          *string              `json:"DeliverToPlant"`
	DeliverToPlantStorageLocation           *string              `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlant                        *string              `json:"DeliverFromPlant"`
	DeliverFromPlantStorageLocation         *string              `json:"DeliverFromPlantStorageLocation"`
	ProductionPlantBusinessPartner          *int                 `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                         *string              `json:"ProductionPlant"`
	ProductionPlantStorageLocation          *string              `json:"ProductionPlantStorageLocation"`
	ServicesRenderedDate                    *string              `json:"ServicesRenderedDate"`
	InvoiceQuantity                         *float32             `json:"InvoiceQuantity"`
	InvoiceQuantityUnit                     *string              `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit               *float32             `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                                *string              `json:"BaseUnit"`
	ActualGoodsIssueDate                    *string              `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                    *string              `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                  *string              `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                  *string              `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                         *float32             `json:"ItemGrossWeight"`
	ItemNetWeight                           *float32             `json:"ItemNetWeight"`
	ItemWeightUnit                          *string              `json:"ItemWeightUnit"`
	NetAmount                               *float32             `json:"NetAmount"`
	TaxAmount                               *float32             `json:"TaxAmount"`
	GrossAmount                             *float32             `json:"GrossAmount"`
	GoodsIssueOrReceiptSlipNumber           *string              `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency                     *string              `json:"TransactionCurrency"`
	PricingDate                             *string              `json:"PricingDate"`
	TransactionTaxClassification            *string              `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   *string              `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry *string              `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                *string              `json:"DefinedTaxClassification"`
	Project                                 *string              `json:"Project"`
	OrderID                                 *int                 `json:"OrderID"`
	OrderItem                               *int                 `json:"OrderItem"`
	OrderType                               *string              `json:"OrderType"`
	ContractType                            *string              `json:"ContractType"`
	OrderVaridityStartDate                  *string              `json:"OrderVaridityStartDate"`
	OrderVaridityEndDate                    *string              `json:"OrderVaridityEndDate"`
	InvoicePeriodStartDate                  *string              `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string              `json:"InvoicePeriodEndDate"`
	DeliveryDocument                        *int                 `json:"DeliveryDocument"`
	DeliveryDocumentItem                    *int                 `json:"DeliveryDocumentItem"`
	OriginDocument                          *int                 `json:"OriginDocument"`
	OriginDocumentItem                      *int                 `json:"OriginDocumentItem"`
	ReferenceDocument                       *int                 `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int                 `json:"ReferenceDocumentItem"`
	ExternalReferenceDocument               *string              `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem           *string              `json:"ExternalReferenceDocumentItem"`
	TaxCode                                 *string              `json:"TaxCode"`
	TaxRate                                 *float32             `json:"TaxRate"`
	CountryOfOrigin                         *string              `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 *string              `json:"CountryOfOriginLanguage"`
	ItemPaymentRequisitionIsCreated         *bool                `json:"ItemPaymentRequisitionIsCreated"`
	ItemIsCleared                           *bool                `json:"ItemIsCleared"`
	ItemPaymentBlockStatus                  *bool                `json:"ItemPaymentBlockStatus"`
	IsCancelled                             *bool                `json:"IsCancelled"`
	ItemPricingElement                      []ItemPricingElement `json:"ItemPricingElement"`
}

type ItemPricingElement struct {
	InvoiceDocument            int      `json:"InvoiceDocument"`
	InvoiceDocumentItem        int      `json:"InvoiceDocumentItem"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}
