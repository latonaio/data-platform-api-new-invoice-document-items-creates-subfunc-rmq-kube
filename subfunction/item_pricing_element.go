package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) OrdersItemPricingElement(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ItemPricingElement, error) {
	args := make([]interface{}, 0)

	ordersItem := psdc.OrdersItem
	repeat := strings.Repeat("(?, ?),", len(ordersItem)-1) + "(?, ?)"
	for _, v := range ordersItem {
		args = append(args, v.OrderID, v.OrderItem)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, PricingProcedureCounter, ConditionRecord, ConditionSequentialNumber, 
		ConditionType, PricingDate, ConditionRateValue, ConditionCurrency, ConditionQuantity, ConditionQuantityUnit, 
		ConditionAmount, ConditionIsManuallyChanged
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_pricing_element_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToItemPricingElement(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentItemPricingElement(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.ItemPricingElement, error) {
	args := make([]interface{}, 0)

	deliveryDocumentItem := psdc.DeliveryDocumentItemData
	repeat := strings.Repeat("(?, ?),", len(deliveryDocumentItem)-1) + "(?, ?)"
	for _, v := range deliveryDocumentItem {
		args = append(args, v.OrderID, v.OrderItem)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, PricingProcedureCounter, ConditionRecord, ConditionSequentialNumber, 
		ConditionType, PricingDate, ConditionRateValue, ConditionCurrency, ConditionQuantity, ConditionQuantityUnit, 
		ConditionAmount, ConditionIsManuallyChanged
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_pricing_element_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToItemPricingElement(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
