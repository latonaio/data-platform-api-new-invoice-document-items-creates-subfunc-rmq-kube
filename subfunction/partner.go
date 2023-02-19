package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) OrdersPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersPartner, error) {
	args := make([]interface{}, 0)

	ordersHeader := psdc.OrdersHeader
	repeat := strings.Repeat("?,", len(ordersHeader)-1) + "?"
	for _, v := range ordersHeader {
		args = append(args, v.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName,
		Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_partner_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersPartner(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentPartner, error) {
	args := make([]interface{}, 0)

	deliveryDocumentHeader := psdc.DeliveryDocumentHeader
	repeat := strings.Repeat("?,", len(deliveryDocumentHeader)-1) + "?"
	for _, v := range deliveryDocumentHeader {
		args = append(args, v.DeliveryDocument)
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName,
		Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_partner_data
		WHERE DeliveryDocument IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentPartner(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
