package api_processing_data_formatter

type SDC struct {
	MetaData                   *MetaData                     `json:"MetaData"`
	ProcessType                *ProcessType                  `json:"ProcessType"`
	ReferenceType              *ReferenceType                `json:"ReferenceType"`
	OrderID                    []*OrderID                    `json:"OrderID"`
	OrderItem                  []*OrderItem                  `json:"OrderItem"`
	DeliveryDocumentHeader     []*DeliveryDocumentHeader     `json:"DeliveryDocumentHeader"`
	DeliveryDocumentItem       []*DeliveryDocumentItem       `json:"DeliveryDocumentItem"`
	OrdersHeader               []*OrdersHeader               `json:"OrdersHeader"`
	OrdersItem                 []*OrdersItem                 `json:"OrdersItem"`
	OrdersPartner              []*OrdersPartner              `json:"OrdersPartner"`
	ItemPricingElement         []*ItemPricingElement         `json:"ItemPricingElement"`
	DeliveryDocumentHeaderData []*DeliveryDocumentHeaderData `json:"DeliveryDocumentHeaderData"`
	DeliveryDocumentItemData   []*DeliveryDocumentItemData   `json:"DeliveryDocumentItemData"`
	DeliveryDocumentPartner    []*DeliveryDocumentPartner    `json:"DeliveryDocumentPartner"`
	CalculateInvoiceDocument   []*CalculateInvoiceDocument   `json:"CalculateInvoiceDocument"`
	InvoiceDocumentItem        []*InvoiceDocumentItem        `json:"InvoiceDocumentItem"`
	Address                    []*Address                    `json:"Address"`
	AddressMaster              []*AddressMaster              `json:"AddressMaster"`
	CreationDateItem           *CreationDate                 `json:"CreationDateItem"`
	LastChangeDateItem         *LastChangeDate               `json:"LastChangeDateItem"`
	CreationTimeItem           *CreationTime                 `json:"CreationTimeItem"`
	LastChangeTimeItem         *LastChangeTime               `json:"LastChangeTimeItem"`
}

// Initializer
type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type ProcessType struct {
	BulkProcess       bool `json:"BulkProcess"`
	IndividualProcess bool `json:"IndividualProcess"`
	ArraySpec         bool `json:"ArraySpec"`
	RangeSpec         bool `json:"RangeSpec"`
}

type ReferenceType struct {
	OrderID          bool `json:"OrderID"`
	DeliveryDocument bool `json:"DeliveryDocument"`
}

type OrderIDKey struct {
	BillFromParty                   []*int `json:"BillFromParty"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToParty                     []*int `json:"BillToParty"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	HeaderCompleteDeliveryIsDefined bool   `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        bool   `json:"HeaderBillingBlockStatus"`
	IsCancelled                     bool   `json:"IsCancelled"`
	IsMarkedForDeletion             bool   `json:"IsMarkedForDeletion"`
	ReferenceDocument               int    `json:"ReferenceDocument"`
}

type OrderID struct {
	OrderID                         int    `json:"OrderID"`
	BillFromParty                   *int   `json:"BillFromParty"`
	BillToParty                     *int   `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
	IsCancelled                     *bool  `json:"IsCancelled"`
	IsMarkedForDeletion             *bool  `json:"IsMarkedForDeletion"`
}

type OrderItemKey struct {
	OrderID                       []int  `json:"OrderID"`
	ItemCompleteDeliveryIsDefined bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            string `json:"ItemDeliveryStatus"`
	ItemBillingStatus             string `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        bool   `json:"ItemBillingBlockStatus"`
	IsCancelled                   bool   `json:"IsCancelled"`
	IsMarkedForDeletion           bool   `json:"IsMarkedForDeletion"`
}

type OrderItem struct {
	OrderID                       int     `json:"OrderID"`
	OrderItem                     int     `json:"OrderItem"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	ItemDeliveryStatus            *string `json:"ItemDeliveryStatus"`
	ItemBillingStatus             *string `json:"ItemBillingStatus"`
	ItemBillingBlockStatus        *bool   `json:"ItemBillingBlockStatus"`
	IsCancelled                   *bool   `json:"IsCancelled"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentHeaderKey struct {
	BillFromParty                   []*int `json:"BillFromParty"`
	BillFromPartyFrom               *int   `json:"BillFromPartyFrom"`
	BillFromPartyTo                 *int   `json:"BillFromPartyTo"`
	BillToParty                     []*int `json:"BillToParty"`
	BillToPartyFrom                 *int   `json:"BillToPartyFrom"`
	BillToPartyTo                   *int   `json:"BillToPartyTo"`
	HeaderCompleteDeliveryIsDefined bool   `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        bool   `json:"HeaderBillingBlockStatus"`
	IsCancelled                     bool   `json:"IsCancelled"`
	IsMarkedForDeletion             bool   `json:"IsMarkedForDeletion"`
	ReferenceDocument               int    `json:"ReferenceDocument"`
}

type DeliveryDocumentHeader struct {
	DeliveryDocument                int    `json:"DeliveryDocument"`
	BillFromParty                   *int   `json:"BillFromParty"`
	BillToParty                     *int   `json:"BillToParty"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	HeaderDeliveryStatus            string `json:"HeaderDeliveryStatus"`
	HeaderBillingStatus             string `json:"HeaderBillingStatus"`
	HeaderBillingBlockStatus        *bool  `json:"HeaderBillingBlockStatus"`
	IsCancelled                     *bool  `json:"IsCancelled"`
	IsMarkedForDeletion             *bool  `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentItemKey struct {
	DeliveryDocument              []int     `json:"DeliveryDocument"`
	ConfirmedDeliveryDate         []*string `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryDateFrom     *string   `json:"BillFromPartyFrom"`
	ConfirmedDeliveryDateTo       *string   `json:"BillFromPartyTo"`
	ActualGoodsIssueDate          []*string `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueDateFrom      *string   `json:"BillToPartyFrom"`
	ActualGoodsIssueDateTo        *string   `json:"BillToPartyTo"`
	ItemCompleteDeliveryIsDefined bool      `json:"ItemCompleteDeliveryIsDefined"`
	// ItemDeliveryStatus            string    `json:"ItemDeliveryStatus"`
	ItemBillingStatus      string `json:"ItemBillingStatus"`
	ItemBillingBlockStatus bool   `json:"ItemBillingBlockStatus"`
	IsCancelled            bool   `json:"IsCancelled"`
	IsMarkedForDeletion    bool   `json:"IsMarkedForDeletion"`
}

type DeliveryDocumentItem struct {
	DeliveryDocument              int     `json:"DeliveryDocument"`
	DeliveryDocumentItem          int     `json:"DeliveryDocumentItem"`
	ConfirmedDeliveryDate         *string `json:"ConfirmedDeliveryDate"`
	ActualGoodsIssueDate          *string `json:"ActualGoodsIssueDate"`
	ItemCompleteDeliveryIsDefined *bool   `json:"ItemCompleteDeliveryIsDefined"`
	// ItemDeliveryStatus            string  `json:"ItemDeliveryStatus"`
	ItemBillingStatus      *string `json:"ItemBillingStatus"`
	ItemBillingBlockStatus *bool   `json:"ItemBillingBlockStatus"`
	IsCancelled            *bool   `json:"IsCancelled"`
	IsMarkedForDeletion    *bool   `json:"IsMarkedForDeletion"`
}

// Orders
type OrdersHeader struct {
	OrderID                          int     `json:"OrderID"`
	OrderType                        string  `json:"OrderType"`
	SupplyChainRelationshipID        int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipBillingID *int    `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID *int    `json:"SupplyChainRelationshipPaymentID"`
	Buyer                            int     `json:"Buyer"`
	Seller                           int     `json:"Seller"`
	BillToParty                      *int    `json:"BillToParty"`
	BillFromParty                    *int    `json:"BillFromParty"`
	BillToCountry                    *string `json:"BillToCountry"`
	BillFromCountry                  *string `json:"BillFromCountry"`
	Payer                            *int    `json:"Payer"`
	Payee                            *int    `json:"Payee"`
	ContractType                     *string `json:"ContractType"`
	OrderValidityStartDate           *string `json:"OrderVaridityStartDate"`
	OrderValidityEndDate             *string `json:"OrderValidityEndDate"`
	InvoicePeriodStartDate           *string `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate             *string `json:"InvoicePeriodEndDate"`
	TotalNetAmount                   float32 `json:"TotalNetAmount"`
	TotalTaxAmount                   float32 `json:"TotalTaxAmount"`
	TotalGrossAmount                 float32 `json:"TotalGrossAmount"`
	TransactionCurrency              string  `json:"TransactionCurrency"`
	PricingDate                      string  `json:"PricingDate"`
	Incoterms                        *string `json:"Incoterms"`
	PaymentTerms                     string  `json:"PaymentTerms"`
	PaymentMethod                    string  `json:"PaymentMethod"`
	IsExportImport                   *bool   `json:"IsExportImport"`
}

type OrdersItem struct {
	OrderID                                 int      `json:"OrderID"`
	OrderItem                               int      `json:"OrderItem"`
	OrderItemCategory                       string   `json:"OrderItemCategory"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	OrderItemText                           string   `json:"OrderItemText"`
	OrderItemTextByBuyer                    string   `json:"OrderItemTextByBuyer"`
	OrderItemTextBySeller                   string   `json:"OrderItemTextBySeller"`
	Product                                 string   `json:"Product"`
	ProductStandardID                       string   `json:"ProductStandardID"`
	ProductGroup                            *string  `json:"ProductGroup"`
	BaseUnit                                string   `json:"BaseUnit"`
	PricingDate                             string   `json:"PricingDate"`
	DeliveryUnit                            string   `json:"DeliveryUnit"`
	OrderQuantityInBaseUnit                 float32  `json:"OrderQuantityInBaseUnit"`
	OrderQuantityInDeliveryUnit             float32  `json:"OrderQuantityInDeliveryUnit"`
	NetAmount                               *float32 `json:"NetAmount"`
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	Incoterms                               *string  `json:"Incoterms"`
	TransactionTaxClassification            string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassification                string   `json:"DefinedTaxClassification"`
	PaymentTerms                            string   `json:"PaymentTerms"`
	PaymentMethod                           string   `json:"PaymentMethod"`
	Project                                 *string  `json:"Project"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
}

type OrdersPartner struct {
	OrderID                 int     `json:"OrderID"`
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

type ItemPricingElement struct {
	OrderID                    int      `json:"OrderID"`
	OrderItem                  int      `json:"OrderItem"`
	PricingProcedureCounter    int      `json:"PricingProcedureCounter"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

// DeliveryDocument
type DeliveryDocumentHeaderData struct {
	DeliveryDocument                       int     `json:"DeliveryDocument"`
	SupplyChainRelationshipID              int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID      int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID       *int    `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID       *int    `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                  int     `json:"Buyer"`
	Seller                                 int     `json:"Seller"`
	DeliverToParty                         int     `json:"DeliverToParty"`
	DeliverFromParty                       int     `json:"DeliverFromParty"`
	DeliverToPlant                         string  `json:"DeliverToPlant"`
	DeliverFromPlant                       string  `json:"DeliverFromPlant"`
	BillToParty                            *int    `json:"BillToParty"`
	BillFromParty                          *int    `json:"BillFromParty"`
	BillToCountry                          *string `json:"BillToCountry"`
	BillFromCountry                        *string `json:"BillFromCountry"`
	Payer                                  *int    `json:"Payer"`
	Payee                                  *int    `json:"Payee"`
	IsExportImport                         *bool   `json:"IsExportImport"`
	OrderID                                *int    `json:"OrderID"`
	OrderItem                              *int    `json:"OrderItem"`
	ContractType                           *string `json:"ContractType"`
	OrderValidityStartDate                 *string `json:"OrderValidityStartDate"`
	OrderValidityEndDate                   *string `json:"OrderValidityEndDate"`
	GoodsIssueOrReceiptSlipNumber          *string `json:"GoodsIssueOrReceiptSlipNumber"`
	Incoterms                              *string `json:"Incoterms"`
	TransactionCurrency                    *string `json:"TransactionCurrency"`
}

type DeliveryDocumentItemData struct {
	DeliveryDocument                        int      `json:"DeliveryDocument"`
	DeliveryDocumentItem                    int      `json:"DeliveryDocumentItem"`
	DeliveryDocumentItemCategory            *string  `json:"DeliveryDocumentItemCategory"`
	SupplyChainRelationshipID               int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID       int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID  int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipBillingID        *int     `json:"SupplyChainRelationshipBillingID"`
	SupplyChainRelationshipPaymentID        *int     `json:"SupplyChainRelationshipPaymentID"`
	Buyer                                   int      `json:"Buyer"`
	Seller                                  int      `json:"Seller"`
	DeliverToParty                          int      `json:"DeliverToParty"`
	DeliverFromParty                        int      `json:"DeliverFromParty"`
	DeliverToPlant                          string   `json:"DeliverToPlant"`
	DeliverFromPlant                        string   `json:"DeliverFromPlant"`
	BillToParty                             *int     `json:"BillToParty"`
	BillFromParty                           *int     `json:"BillFromParty"`
	BillToCountry                           *string  `json:"BillToCountry"`
	BillFromCountry                         *string  `json:"BillFromCountry"`
	Payer                                   *int     `json:"Payer"`
	Payee                                   *int     `json:"Payee"`
	DeliverToPlantStorageLocation           *string  `json:"DeliverToPlantStorageLocation"`
	DeliverFromPlantStorageLocation         *string  `json:"DeliverFromPlantStorageLocation"`
	ProductionPlantBusinessPartner          *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                         *string  `json:"ProductionPlant"`
	ProductionPlantStorageLocation          *string  `json:"ProductionPlantStorageLocation"`
	DeliveryDocumentItemText                *string  `json:"DeliveryDocumentItemText"`
	DeliveryDocumentItemTextByBuyer         string   `json:"DeliveryDocumentItemTextByBuyer"`
	DeliveryDocumentItemTextBySeller        string   `json:"DeliveryDocumentItemTextBySeller"`
	Product                                 *string  `json:"Product"`
	ProductStandardID                       *string  `json:"ProductStandardID"`
	ProductGroup                            *string  `json:"ProductGroup"`
	BaseUnit                                *string  `json:"BaseUnit"`
	DeliveryUnit                            *string  `json:"DeliveryUnit"`
	ActualGoodsIssueDate                    *string  `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime                    *string  `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate                  *string  `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime                  *string  `json:"ActualGoodsReceiptTime"`
	ActualGoodsIssueQuantity                *float32 `json:"ActualGoodsIssueQuantity"`
	ActualGoodsIssueQtyInBaseUnit           *float32 `json:"ActualGoodsIssueQtyInBaseUnit"`
	ActualGoodsReceiptQuantity              *float32 `json:"ActualGoodsReceiptQuantity"`
	ActualGoodsReceiptQtyInBaseUnit         *float32 `json:"ActualGoodsReceiptQtyInBaseUnit"`
	ItemGrossWeight                         *float32 `json:"ItemGrossWeight"`
	ItemNetWeight                           *float32 `json:"ItemNetWeight"`
	ItemWeightUnit                          *string  `json:"ItemWeightUnit"`
	NetAmount                               *float32 `json:"NetAmount"`
	TaxAmount                               *float32 `json:"TaxAmount"`
	GrossAmount                             *float32 `json:"GrossAmount"`
	OrderID                                 *int     `json:"OrderID"`
	OrderItem                               *int     `json:"OrderItem"`
	OrderType                               *string  `json:"OrderType"`
	ContractType                            *string  `json:"ContractType"`
	OrderValidityStartDate                  *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate                    *string  `json:"OrderValidityEndDate"`
	PaymentTerms                            *string  `json:"PaymentTerms"`
	PaymentMethod                           *string  `json:"PaymentMethod"`
	InvoicePeriodStartDate                  *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate                    *string  `json:"InvoicePeriodEndDate"`
	Project                                 *string  `json:"Project"`
	ReferenceDocument                       *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem                   *int     `json:"ReferenceDocumentItem"`
	TransactionTaxClassification            string   `json:"TransactionTaxClassification"`
	ProductTaxClassificationBillToCountry   string   `json:"ProductTaxClassificationBillToCountry"`
	ProductTaxClassificationBillFromCountry string   `json:"ProductTaxClassificationBillFromCountry"`
	DefinedTaxClassifications               string   `json:"DefinedTaxClassification"`
	TaxCode                                 *string  `json:"TaxCode"`
	TaxRate                                 *float32 `json:"TaxRate"`
	CountryOfOrigin                         *string  `json:"CountryOfOrigin"`
	CountryOfOriginLanguage                 *string  `json:"CountryOfOriginLanguage"`
}

type DeliveryDocumentPartner struct {
	DeliveryDocument        int     `json:"DeliveryDocument"`
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

// Header
type CalculateInvoiceDocument struct {
	InvoiceDocumentLatestNumber *int `json:"InvoiceDocumentLatestNumber"`
	InvoiceDocument             int  `json:"InvoiceDocument"`
	OrderID                     int  `json:"OrderID"`
	OrderItem                   int  `json:"OrderItem"`
	DeliveryDocument            int  `json:"DeliveryDocument"`
	DeliveryDocumentItem        int  `json:"DeliveryDocumentItem"`
	BillFromParty               int  `json:"BillFromParty"`
	BillToParty                 int  `json:"BillToParty"`
}

type BillParty struct {
	BillFromParty        int `json:"BillFromParty"`
	BillToParty          int `json:"BillToParty"`
	OrderID              int `json:"OrderID"`
	OrderItem            int `json:"OrderItem"`
	DeliveryDocument     int `json:"DeliveryDocument"`
	DeliveryDocumentItem int `json:"DeliveryDocumentItem"`
}

// Item
type InvoiceDocumentItem struct {
	OrderID                   int `json:"OrderID"`
	OrderItem                 int `json:"OrderItem"`
	DeliveryDocument          int `json:"DeliveryDocument"`
	DeliveryDocumentItem      int `json:"DeliveryDocumentItem"`
	InvoiceDocumentItemNumber int `json:"InvoiceDocumentItemNumber"`
}

// Address
type Address struct {
	OrderID          int     `json:"OrderID"`
	DeliveryDocument int     `json:"DeliveryDocument"`
	AddressID        int     `json:"AddressID"`
	PostalCode       *string `json:"PostalCode"`
	LocalRegion      *string `json:"LocalRegion"`
	Country          *string `json:"Country"`
	District         *string `json:"District"`
	StreetName       *string `json:"StreetName"`
	CityName         *string `json:"CityName"`
	Building         *string `json:"Building"`
	Floor            *int    `json:"Floor"`
	Room             *int    `json:"Room"`
}

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

type CalculateAddressIDKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateAddressIDQueryGets struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
	LatestNumber             *int   `json:"LatestNumber"`
}

type CalculateAddressID struct {
	AddressIDLatestNumber *int `json:"AddressIDLatestNumber"`
	AddressID             int  `json:"AddressID"`
}

// 日付等の処理
type CreationDate struct {
	CreationDate string `json:"CreationDate"`
}

type LastChangeDate struct {
	LastChangeDate string `json:"LastChangeDate"`
}

type CreationTime struct {
	CreationTime string `json:"CreationTime"`
}

type LastChangeTime struct {
	LastChangeTime string `json:"LastChangeTime"`
}
