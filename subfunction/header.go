package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
	"time"

	"golang.org/x/xerrors"
)

func (f *SubFunction) OrdersHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersHeader, error) {
	args := make([]interface{}, 0)

	orderID := psdc.OrderID
	repeat := strings.Repeat("?,", len(orderID)-1) + "?"
	for _, tag := range orderID {
		args = append(args, tag.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderType, SupplyChainRelationshipID, SupplyChainRelationshipBillingID,
		SupplyChainRelationshipPaymentID, Buyer, Seller, BillToParty, BillFromParty, BillToCountry, BillFromCountry,
		Payer, Payee, ContractType, OrderValidityStartDate, OrderValidityEndDate, InvoicePeriodStartDate,
		InvoicePeriodEndDate, TotalNetAmount, TotalTaxAmount, TotalGrossAmount, TransactionCurrency,
		PricingDate, Incoterms, PaymentTerms, PaymentMethod, IsExportImport
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentHeaderData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentHeaderData, error) {
	args := make([]interface{}, 0)

	deliveryDocument := psdc.DeliveryDocumentItem
	repeat := strings.Repeat("?,", len(deliveryDocument)-1) + "?"
	for _, tag := range deliveryDocument {
		args = append(args, tag.DeliveryDocument)
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, SupplyChainRelationshipID, SupplyChainRelationshipDeliveryID, SupplyChainRelationshipDeliveryPlantID,
		SupplyChainRelationshipBillingID, SupplyChainRelationshipPaymentID, Buyer, Seller, DeliverToParty, DeliverFromParty, 
		DeliverToPlant, DeliverFromPlant, BillToParty, BillFromParty, BillToCountry, BillFromCountry, Payer, Payee, IsExportImport,
		OrderID, OrderItem, ContractType, OrderValidityStartDate, OrderValidityEndDate, GoodsIssueOrReceiptSlipNumber,
		Incoterms, TransactionCurrency
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE DeliveryDocument IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentHeaderData(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CalculateInvoiceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.CalculateInvoiceDocument, error) {
	metaData := psdc.MetaData
	dataKey := psdc.ConvertToInvoiceDocumentHeaderKey()

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataQueryGets, err := psdc.ConvertToInvoiceDocumentHeaderQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	if dataQueryGets.InvoiceDocumentLatestNumber == nil {
		return nil, xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルのLatestNumberがNULLです。")
	}

	billParties := make([]api_processing_data_formatter.BillParty, 0)
	if psdc.ReferenceType.OrderID {
		for _, ordersHeader := range psdc.OrdersHeader {
			billFromParty := ordersHeader.BillFromParty
			billToParty := ordersHeader.BillToParty
			orderID := ordersHeader.OrderID

			if billFromParty == nil || billToParty == nil {
				continue
			}

			if billPartyContain(billParties, *billFromParty, *billToParty) {
				continue
			}

			billParties = append(billParties, api_processing_data_formatter.BillParty{
				BillFromParty: *billFromParty,
				BillToParty:   *billToParty,
				OrderID:       orderID,
			})
		}
	} else if psdc.ReferenceType.DeliveryDocument {
		for _, deliveryDocumentItemData := range psdc.DeliveryDocumentItemData {
			billFromParty := deliveryDocumentItemData.BillFromParty
			billToParty := deliveryDocumentItemData.BillToParty
			deliveryDocument := deliveryDocumentItemData.DeliveryDocument
			deliveryDocumentItem := deliveryDocumentItemData.DeliveryDocumentItem

			if billFromParty == nil || billToParty == nil {
				continue
			}

			if billPartyContain(billParties, *billFromParty, *billToParty) {
				continue
			}

			billParties = append(billParties, api_processing_data_formatter.BillParty{
				BillFromParty:        *billFromParty,
				BillToParty:          *billToParty,
				DeliveryDocument:     deliveryDocument,
				DeliveryDocumentItem: deliveryDocumentItem,
			})
		}
	}

	data := make([]*api_processing_data_formatter.CalculateInvoiceDocument, 0)
	for i, billParty := range billParties {
		invoiceDocumentLatestNumber := dataQueryGets.InvoiceDocumentLatestNumber
		invoiceDocument := *dataQueryGets.InvoiceDocumentLatestNumber + i + 1
		billFromParty := billParty.BillFromParty
		billToParty := billParty.BillToParty
		orderID := billParty.OrderID
		deliveryDocument := billParty.DeliveryDocument
		deliveryDocumentItem := billParty.DeliveryDocumentItem

		datum := psdc.ConvertToCalculateInvoiceDocument(invoiceDocumentLatestNumber, invoiceDocument, orderID, deliveryDocument, deliveryDocumentItem, billFromParty, billToParty)
		data = append(data, datum)
	}

	return data, err
}

func billPartyContain(billParties []api_processing_data_formatter.BillParty, billFromParty, billToParty int) bool {
	for _, billParty := range billParties {
		if billFromParty == billParty.BillFromParty && billToParty == billParty.BillToParty {
			return true
		}
	}
	return false
}

func getSystemTime() string {
	day := time.Now()
	return day.Format("15:04:05")
}
