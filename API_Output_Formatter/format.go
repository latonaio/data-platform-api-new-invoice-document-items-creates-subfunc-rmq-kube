package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"encoding/json"
	"reflect"

	"golang.org/x/xerrors"
)

func ConvertToItem(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Item, error) {
	var err error
	items := make([]*Item, 0)

	ordersHeaderMap := StructArrayToMap(psdc.OrdersHeader, "OrderID")
	deliveryDocumentHeaderMap := StructArrayToMap(psdc.DeliveryDocumentHeader, "DeliveryDocument")

	processType := psdc.ProcessType
	referenceType := psdc.ReferenceType
	if referenceType.OrderID {
		if processType.BulkProcess {
			for i := range psdc.InvoiceDocumentItem {
				item := &Item{}
				inputItem := sdc.Header.Item[0]

				orderID := psdc.InvoiceDocumentItem[i].OrderID
				orderItem := psdc.InvoiceDocumentItem[i].OrderItem

				ordersItemIdx := -1
				for j, ordersItem := range psdc.OrdersItem {
					if ordersItem.OrderID == orderID && ordersItem.OrderItem == orderItem {
						ordersItemIdx = j
						break
					}
				}
				if ordersItemIdx == -1 {
					continue
				}

				var billFromParty, billToParty int
				for _, ordersHeader := range psdc.OrdersHeader {
					if ordersHeader.OrderID != orderID {
						continue
					}

					if ordersHeader.BillFromParty == nil || ordersHeader.BillToParty == nil {
						continue
					}

					billFromParty = *ordersHeader.BillFromParty
					billToParty = *ordersHeader.BillToParty
					break
				}

				// 入力ファイル
				item, err = jsonTypeConversion(item, inputItem)
				if err != nil {
					return nil, err
				}

				// 1-2
				item, err = jsonTypeConversion(item, psdc.OrdersItem[ordersItemIdx])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				// 1-1
				if _, ok := ordersHeaderMap[orderID]; !ok {
					continue
				}
				item, err = jsonTypeConversion(item, ordersHeaderMap[orderID])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				invoiceDocumentIdx := -1
				for j, invoiceDocument := range psdc.CalculateInvoiceDocument {
					if invoiceDocument.BillFromParty == billFromParty && invoiceDocument.BillToParty == billToParty {
						invoiceDocumentIdx = j
						break
					}
				}
				if invoiceDocumentIdx == -1 {
					continue
				}

				item.InvoiceDocument = psdc.CalculateInvoiceDocument[invoiceDocumentIdx].InvoiceDocument
				item.InvoiceDocumentItem = psdc.InvoiceDocumentItem[i].InvoiceDocumentItemNumber
				item.InvoiceDocumentItemCategory = &psdc.OrdersItem[ordersItemIdx].OrderItemCategory
				item.InvoiceDocumentItemText = &psdc.OrdersItem[ordersItemIdx].OrderItemText
				item.InvoiceDocumentItemTextByBuyer = psdc.OrdersItem[ordersItemIdx].OrderItemTextByBuyer
				item.InvoiceDocumentItemTextBySeller = psdc.OrdersItem[ordersItemIdx].OrderItemTextBySeller

				item.CreationDate = psdc.CreationDateItem.CreationDate
				item.CreationTime = psdc.CreationTimeItem.CreationTime
				item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
				item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime
				item.ItemBillingIsConfirmed = getBoolPtr(false)

				item.InvoiceQuantity = &psdc.OrdersItem[ordersItemIdx].OrderQuantityInDeliveryUnit
				item.InvoiceQuantityUnit = &psdc.OrdersItem[ordersItemIdx].DeliveryUnit
				item.InvoiceQuantityInBaseUnit = &psdc.OrdersItem[ordersItemIdx].OrderQuantityInBaseUnit

				item.OriginDocument = psdc.OrdersItem[ordersItemIdx].ReferenceDocument
				item.OriginDocumentItem = psdc.OrdersItem[ordersItemIdx].ReferenceDocumentItem

				item.ItemPaymentRequisitionIsCreated = getBoolPtr(false)
				item.ItemIsCleared = getBoolPtr(false)
				item.ItemPaymentBlockStatus = getBoolPtr(false)
				item.IsCancelled = getBoolPtr(false)

				items = append(items, item)
			}
		} else if processType.IndividualProcess {
			for i := range psdc.InvoiceDocumentItem {
				item := &Item{}
				inputItem := sdc.Header.Item[0]

				orderID := psdc.InvoiceDocumentItem[i].OrderID
				orderItem := psdc.InvoiceDocumentItem[i].OrderItem

				ordersItemIdx := -1
				for j, ordersItem := range psdc.OrdersItem {
					if ordersItem.OrderID == orderID && ordersItem.OrderItem == orderItem {
						ordersItemIdx = j
						break
					}
				}
				if ordersItemIdx == -1 {
					continue
				}

				// 入力ファイル
				item, err = jsonTypeConversion(item, inputItem)
				if err != nil {
					return nil, err
				}

				// 1-1
				item, err = jsonTypeConversion(item, psdc.OrdersHeader[0])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				// 1-2
				item, err = jsonTypeConversion(item, psdc.OrdersItem[i])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				item.InvoiceDocument = psdc.CalculateInvoiceDocument[0].InvoiceDocument
				item.InvoiceDocumentItem = psdc.InvoiceDocumentItem[i].InvoiceDocumentItemNumber
				item.InvoiceDocumentItemCategory = &psdc.OrdersItem[i].OrderItemCategory
				item.InvoiceDocumentItemText = &psdc.OrdersItem[ordersItemIdx].OrderItemText
				item.InvoiceDocumentItemTextByBuyer = psdc.OrdersItem[ordersItemIdx].OrderItemTextByBuyer
				item.InvoiceDocumentItemTextBySeller = psdc.OrdersItem[ordersItemIdx].OrderItemTextBySeller

				item.CreationDate = psdc.CreationDateItem.CreationDate
				item.CreationTime = psdc.CreationTimeItem.CreationTime
				item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
				item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime
				item.ItemBillingIsConfirmed = getBoolPtr(false)

				item.InvoiceQuantity = &psdc.OrdersItem[ordersItemIdx].OrderQuantityInDeliveryUnit
				item.InvoiceQuantityUnit = &psdc.OrdersItem[ordersItemIdx].DeliveryUnit
				item.InvoiceQuantityInBaseUnit = &psdc.OrdersItem[ordersItemIdx].OrderQuantityInBaseUnit

				item.OriginDocument = psdc.OrdersItem[ordersItemIdx].ReferenceDocument
				item.OriginDocumentItem = psdc.OrdersItem[ordersItemIdx].ReferenceDocumentItem

				item.ItemPaymentRequisitionIsCreated = getBoolPtr(false)
				item.ItemIsCleared = getBoolPtr(false)
				item.ItemPaymentBlockStatus = getBoolPtr(false)
				item.IsCancelled = getBoolPtr(false)

				items = append(items, item)
			}
		}
	} else if referenceType.DeliveryDocument {
		if processType.BulkProcess {
			for i := range psdc.InvoiceDocumentItem {
				item := &Item{}
				inputItem := sdc.Header.Item[0]

				deliveryDocument := psdc.InvoiceDocumentItem[i].DeliveryDocument
				deliveryDocumentItem := psdc.InvoiceDocumentItem[i].DeliveryDocumentItem

				deliveryDocumentItemIdx := -1
				var billFromParty, billToParty int

				for j, deliveryDocumentItemData := range psdc.DeliveryDocumentItemData {
					if deliveryDocumentItemData.DeliveryDocument == deliveryDocument && deliveryDocumentItemData.DeliveryDocumentItem == deliveryDocumentItem {
						if deliveryDocumentItemData.BillFromParty == nil || deliveryDocumentItemData.BillToParty == nil {
							continue
						}

						billFromParty = *deliveryDocumentItemData.BillFromParty
						billToParty = *deliveryDocumentItemData.BillToParty

						deliveryDocumentItemIdx = j
						break
					}
				}
				if deliveryDocumentItemIdx == -1 {
					continue
				}

				// 入力ファイル
				item, err = jsonTypeConversion(item, inputItem)
				if err != nil {
					return nil, err
				}

				// 1-2
				item, err = jsonTypeConversion(item, psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				// 1-1
				if _, ok := deliveryDocumentHeaderMap[deliveryDocument]; !ok {
					continue
				}
				item, err = jsonTypeConversion(item, deliveryDocumentHeaderMap[deliveryDocument])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				invoiceDocumentIdx := -1
				for j, invoiceDocument := range psdc.CalculateInvoiceDocument {
					if invoiceDocument.BillFromParty == billFromParty && invoiceDocument.BillToParty == billToParty {
						invoiceDocumentIdx = j
						break
					}
				}
				if invoiceDocumentIdx == -1 {
					continue
				}

				item.InvoiceDocument = psdc.CalculateInvoiceDocument[invoiceDocumentIdx].InvoiceDocument
				item.InvoiceDocumentItem = psdc.InvoiceDocumentItem[i].InvoiceDocumentItemNumber
				item.InvoiceDocumentItemCategory = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemCategory
				item.InvoiceDocumentItemText = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemText
				item.InvoiceDocumentItemTextByBuyer = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemTextByBuyer
				item.InvoiceDocumentItemTextBySeller = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemTextBySeller

				item.CreationDate = psdc.CreationDateItem.CreationDate
				item.CreationTime = psdc.CreationTimeItem.CreationTime
				item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
				item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime
				item.ItemBillingIsConfirmed = getBoolPtr(false)

				item.InvoiceQuantity = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ActualGoodsReceiptQuantity
				item.InvoiceQuantityUnit = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryUnit
				item.InvoiceQuantityInBaseUnit = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ActualGoodsReceiptQtyInBaseUnit

				item.OriginDocument = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ReferenceDocument
				item.OriginDocumentItem = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ReferenceDocumentItem

				item.ItemPaymentRequisitionIsCreated = getBoolPtr(false)
				item.ItemIsCleared = getBoolPtr(false)
				item.ItemPaymentBlockStatus = getBoolPtr(false)
				item.IsCancelled = getBoolPtr(false)

				items = append(items, item)
			}
		} else if processType.IndividualProcess {
			for i := range psdc.InvoiceDocumentItem {
				item := &Item{}
				inputItem := sdc.Header.Item[0]

				deliveryDocument := psdc.InvoiceDocumentItem[i].DeliveryDocument
				deliveryDocumentItem := psdc.InvoiceDocumentItem[i].DeliveryDocumentItem

				deliveryDocumentItemIdx := -1
				var billFromParty, billToParty int

				for j, deliveryDocumentItemData := range psdc.DeliveryDocumentItemData {
					if deliveryDocumentItemData.DeliveryDocument == deliveryDocument && deliveryDocumentItemData.DeliveryDocumentItem == deliveryDocumentItem {
						if deliveryDocumentItemData.BillFromParty == nil || deliveryDocumentItemData.BillToParty == nil {
							continue
						}

						billFromParty = *deliveryDocumentItemData.BillFromParty
						billToParty = *deliveryDocumentItemData.BillToParty

						deliveryDocumentItemIdx = j
						break
					}
				}
				if deliveryDocumentItemIdx == -1 {
					continue
				}

				// 入力ファイル
				item, err = jsonTypeConversion(item, inputItem)
				if err != nil {
					return nil, err
				}

				// 1-2
				item, err = jsonTypeConversion(item, psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				// 1-1
				if _, ok := deliveryDocumentHeaderMap[deliveryDocument]; !ok {
					continue
				}
				item, err = jsonTypeConversion(item, deliveryDocumentHeaderMap[deliveryDocument])
				if err != nil {
					return nil, xerrors.Errorf("request create error: %w", err)
				}

				invoiceDocumentIdx := -1
				for j, invoiceDocument := range psdc.CalculateInvoiceDocument {
					if invoiceDocument.BillFromParty == billFromParty && invoiceDocument.BillToParty == billToParty {
						invoiceDocumentIdx = j
						break
					}
				}
				if invoiceDocumentIdx == -1 {
					continue
				}

				item.InvoiceDocument = psdc.CalculateInvoiceDocument[invoiceDocumentIdx].InvoiceDocument
				item.InvoiceDocumentItem = psdc.InvoiceDocumentItem[i].InvoiceDocumentItemNumber
				item.InvoiceDocumentItemCategory = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemCategory
				item.InvoiceDocumentItemText = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemText
				item.InvoiceDocumentItemTextByBuyer = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemTextByBuyer
				item.InvoiceDocumentItemTextBySeller = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryDocumentItemTextBySeller

				item.CreationDate = psdc.CreationDateItem.CreationDate
				item.CreationTime = psdc.CreationTimeItem.CreationTime
				item.LastChangeDate = psdc.LastChangeDateItem.LastChangeDate
				item.LastChangeTime = psdc.LastChangeTimeItem.LastChangeTime
				item.ItemBillingIsConfirmed = getBoolPtr(false)

				item.InvoiceQuantity = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ActualGoodsReceiptQuantity
				item.InvoiceQuantityUnit = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].DeliveryUnit
				item.InvoiceQuantityInBaseUnit = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ActualGoodsReceiptQtyInBaseUnit

				item.OriginDocument = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ReferenceDocument
				item.OriginDocumentItem = psdc.DeliveryDocumentItemData[deliveryDocumentItemIdx].ReferenceDocumentItem

				item.ItemPaymentRequisitionIsCreated = getBoolPtr(false)
				item.ItemIsCleared = getBoolPtr(false)
				item.ItemPaymentBlockStatus = getBoolPtr(false)
				item.IsCancelled = getBoolPtr(false)

				items = append(items, item)
			}
		}
	}

	return items, nil
}

func ConvertToItemPricingElement(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*ItemPricingElement, error) {
	var err error
	itemPricingElements := make([]*ItemPricingElement, 0)

	ordersHeaderMap := StructArrayToMap(psdc.OrdersHeader, "OrderID")

	referenceType := psdc.ReferenceType
	if referenceType.OrderID {
		for _, invoiceDocumentItemPricingElement := range psdc.ItemPricingElement {
			itemPricingElement := &ItemPricingElement{}
			inputItemPricingElement := sdc.Header.Address[0]

			// 入力ファイル
			itemPricingElement, err = jsonTypeConversion(itemPricingElement, inputItemPricingElement)
			if err != nil {
				return nil, err
			}

			orderID := invoiceDocumentItemPricingElement.OrderID
			orderItem := invoiceDocumentItemPricingElement.OrderItem
			if _, ok := ordersHeaderMap[orderID]; !ok {
				continue
			}
			if ordersHeaderMap[orderID].BillFromParty == nil || ordersHeaderMap[orderID].BillToParty == nil {
				continue
			}
			billFromParty := *ordersHeaderMap[orderID].BillFromParty
			billToParty := *ordersHeaderMap[orderID].BillToParty

			invoiceDocumentIdx := -1
			for j, invoiceDocument := range psdc.CalculateInvoiceDocument {
				if invoiceDocument.BillFromParty == billFromParty && invoiceDocument.BillToParty == billToParty {
					invoiceDocumentIdx = j
					break
				}
			}
			if invoiceDocumentIdx == -1 {
				continue
			}

			invoiceDocumentItemIdx := -1
			for j, invoiceDocumentItem := range psdc.InvoiceDocumentItem {
				if invoiceDocumentItem.OrderID == orderID && invoiceDocumentItem.OrderItem == orderItem {
					invoiceDocumentItemIdx = j
					break
				}
			}
			if invoiceDocumentItemIdx == -1 {
				continue
			}

			itemPricingElement.InvoiceDocument = psdc.CalculateInvoiceDocument[invoiceDocumentIdx].InvoiceDocument
			itemPricingElement.InvoiceDocumentItem = psdc.InvoiceDocumentItem[invoiceDocumentItemIdx].InvoiceDocumentItemNumber
			itemPricingElement.PricingProcedureCounter = invoiceDocumentItemPricingElement.PricingProcedureCounter
			itemPricingElement.ConditionRecord = invoiceDocumentItemPricingElement.ConditionRecord
			itemPricingElement.ConditionSequentialNumber = invoiceDocumentItemPricingElement.ConditionSequentialNumber
			itemPricingElement.ConditionType = invoiceDocumentItemPricingElement.ConditionType
			itemPricingElement.PricingDate = invoiceDocumentItemPricingElement.PricingDate
			itemPricingElement.ConditionRateValue = invoiceDocumentItemPricingElement.ConditionRateValue
			itemPricingElement.ConditionCurrency = invoiceDocumentItemPricingElement.ConditionCurrency
			itemPricingElement.ConditionQuantity = invoiceDocumentItemPricingElement.ConditionQuantity
			itemPricingElement.ConditionQuantityUnit = invoiceDocumentItemPricingElement.ConditionQuantityUnit
			// itemPricingElement.TaxCode = invoiceDocumentItemPricingElement.TaxCode
			itemPricingElement.ConditionAmount = invoiceDocumentItemPricingElement.ConditionAmount
			// itemPricingElement.TransactionCurrency = invoiceDocumentItemPricingElement.TransactionCurrency
			itemPricingElement.ConditionIsManuallyChanged = invoiceDocumentItemPricingElement.ConditionIsManuallyChanged

			itemPricingElements = append(itemPricingElements, itemPricingElement)
		}
	} else if referenceType.DeliveryDocument {
		for _, invoiceDocumentItemPricingElement := range psdc.ItemPricingElement {
			itemPricingElement := &ItemPricingElement{}
			inputItemPricingElement := sdc.Header.Address[0]

			// 入力ファイル
			itemPricingElement, err = jsonTypeConversion(itemPricingElement, inputItemPricingElement)
			if err != nil {
				return nil, err
			}

			orderID := invoiceDocumentItemPricingElement.OrderID
			orderItem := invoiceDocumentItemPricingElement.OrderItem
			var deliveryDocument, deliveryDocumentItem int
			var billFromParty, billToParty int
			for _, deliveryDocumentItemData := range psdc.DeliveryDocumentItemData {
				if deliveryDocumentItemData.OrderID == nil || deliveryDocumentItemData.OrderItem == nil {
					continue
				}

				if *deliveryDocumentItemData.OrderID == orderID && *deliveryDocumentItemData.OrderItem == orderItem {
					if deliveryDocumentItemData.BillFromParty == nil || deliveryDocumentItemData.BillToParty == nil {
						continue
					}
					deliveryDocument = deliveryDocumentItemData.DeliveryDocument
					deliveryDocumentItem = deliveryDocumentItemData.DeliveryDocumentItem
					billFromParty = *deliveryDocumentItemData.BillFromParty
					billToParty = *deliveryDocumentItemData.BillToParty
				}
			}

			invoiceDocumentIdx := -1
			for j, invoiceDocument := range psdc.CalculateInvoiceDocument {
				if invoiceDocument.BillFromParty == billFromParty && invoiceDocument.BillToParty == billToParty {
					invoiceDocumentIdx = j
					break
				}
			}
			if invoiceDocumentIdx == -1 {
				continue
			}

			invoiceDocumentItemIdx := -1
			for j, invoiceDocumentItem := range psdc.InvoiceDocumentItem {
				if invoiceDocumentItem.DeliveryDocument == deliveryDocument && invoiceDocumentItem.DeliveryDocumentItem == deliveryDocumentItem {
					invoiceDocumentItemIdx = j
					break
				}
			}
			if invoiceDocumentItemIdx == -1 {
				continue
			}

			itemPricingElement.InvoiceDocument = psdc.CalculateInvoiceDocument[invoiceDocumentIdx].InvoiceDocument
			itemPricingElement.InvoiceDocumentItem = psdc.InvoiceDocumentItem[invoiceDocumentItemIdx].InvoiceDocumentItemNumber
			itemPricingElement.PricingProcedureCounter = invoiceDocumentItemPricingElement.PricingProcedureCounter
			itemPricingElement.ConditionRecord = invoiceDocumentItemPricingElement.ConditionRecord
			itemPricingElement.ConditionSequentialNumber = invoiceDocumentItemPricingElement.ConditionSequentialNumber
			itemPricingElement.ConditionType = invoiceDocumentItemPricingElement.ConditionType
			itemPricingElement.PricingDate = invoiceDocumentItemPricingElement.PricingDate
			itemPricingElement.ConditionRateValue = invoiceDocumentItemPricingElement.ConditionRateValue
			itemPricingElement.ConditionCurrency = invoiceDocumentItemPricingElement.ConditionCurrency
			itemPricingElement.ConditionQuantity = invoiceDocumentItemPricingElement.ConditionQuantity
			itemPricingElement.ConditionQuantityUnit = invoiceDocumentItemPricingElement.ConditionQuantityUnit
			// itemPricingElement.TaxCode = invoiceDocumentItemPricingElement.TaxCode
			itemPricingElement.ConditionAmount = invoiceDocumentItemPricingElement.ConditionAmount
			// itemPricingElement.TransactionCurrency = invoiceDocumentItemPricingElement.TransactionCurrency
			itemPricingElement.ConditionIsManuallyChanged = invoiceDocumentItemPricingElement.ConditionIsManuallyChanged

			itemPricingElements = append(itemPricingElements, itemPricingElement)
		}
	}

	return itemPricingElements, nil
}

func ConvertToPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Partner, error) {
	var err error
	partners := make([]*Partner, 0)

	referenceType := psdc.ReferenceType
	for _, invoiceDocument := range psdc.CalculateInvoiceDocument {
		if referenceType.OrderID {
			for _, invoiceDocumentPartner := range psdc.OrdersPartner {
				partner := &Partner{}
				inputPartner := sdc.Header.Partner[0]

				// 入力ファイル
				partner, err = jsonTypeConversion(partner, inputPartner)
				if err != nil {
					return nil, err
				}

				partner.InvoiceDocument = invoiceDocument.InvoiceDocument
				partner.PartnerFunction = invoiceDocumentPartner.PartnerFunction
				partner.BusinessPartner = invoiceDocumentPartner.BusinessPartner
				partner.BusinessPartnerFullName = invoiceDocumentPartner.BusinessPartnerFullName
				partner.BusinessPartnerName = invoiceDocumentPartner.BusinessPartnerName
				partner.Country = invoiceDocumentPartner.Country
				partner.Language = invoiceDocumentPartner.Language
				partner.Currency = invoiceDocumentPartner.Currency
				partner.ExternalDocumentID = invoiceDocumentPartner.ExternalDocumentID
				partner.AddressID = invoiceDocumentPartner.AddressID

				partners = append(partners, partner)
			}
		} else if referenceType.DeliveryDocument {
			for _, invoiceDocumentPartner := range psdc.DeliveryDocumentPartner {
				partner := &Partner{}
				inputPartner := sdc.Header.Partner[0]

				// 入力ファイル
				partner, err = jsonTypeConversion(partner, inputPartner)
				if err != nil {
					return nil, err
				}

				partner.InvoiceDocument = invoiceDocument.InvoiceDocument
				partner.PartnerFunction = invoiceDocumentPartner.PartnerFunction
				partner.BusinessPartner = invoiceDocumentPartner.BusinessPartner
				partner.BusinessPartnerFullName = invoiceDocumentPartner.BusinessPartnerFullName
				partner.BusinessPartnerName = invoiceDocumentPartner.BusinessPartnerName
				partner.Country = invoiceDocumentPartner.Country
				partner.Language = invoiceDocumentPartner.Language
				partner.Currency = invoiceDocumentPartner.Currency
				partner.ExternalDocumentID = invoiceDocumentPartner.ExternalDocumentID
				partner.AddressID = invoiceDocumentPartner.AddressID

				partners = append(partners, partner)
			}
		}
	}

	return partners, nil
}

func ConvertToAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*Address, error) {
	var err error
	addresses := make([]*Address, 0)

	for _, invoiceDocument := range psdc.CalculateInvoiceDocument {
		for _, invoiceDocumentAddress := range psdc.Address {
			address := &Address{}
			inputAddress := sdc.Header.Address[0]

			// 入力ファイル
			address, err = jsonTypeConversion(address, inputAddress)
			if err != nil {
				return nil, err
			}

			address.InvoiceDocument = invoiceDocument.InvoiceDocument
			address.AddressID = invoiceDocumentAddress.AddressID
			address.PostalCode = invoiceDocumentAddress.PostalCode
			address.LocalRegion = invoiceDocumentAddress.LocalRegion
			address.Country = invoiceDocumentAddress.Country
			address.District = invoiceDocumentAddress.District
			address.StreetName = invoiceDocumentAddress.StreetName
			address.CityName = invoiceDocumentAddress.CityName
			address.Building = invoiceDocumentAddress.Building
			address.Floor = invoiceDocumentAddress.Floor
			address.Room = invoiceDocumentAddress.Room

			addresses = append(addresses, address)
		}
	}

	return addresses, nil
}

func StructArrayToMap[T any](data []T, key string) map[any]T {
	res := make(map[any]T, len(data))

	for _, value := range data {
		m := StructToMap[T](&value, key)
		for k, v := range m {
			res[k] = v
		}
	}

	return res
}

func StructToMap[T any](data interface{}, key string) map[any]T {
	res := make(map[any]T)
	elem := reflect.Indirect(reflect.ValueOf(data).Elem())
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		if field == key {
			rv := reflect.ValueOf(elem.Field(i).Interface())
			if rv.Kind() == reflect.Ptr {
				if rv.IsNil() {
					return nil
				}
			}
			value := reflect.Indirect(elem.Field(i)).Interface()
			var dist T
			res[value], _ = jsonTypeConversion(dist, elem.Interface())
			break
		}
	}

	return res
}

func jsonTypeConversion[T any](dist T, data interface{}) (T, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}
