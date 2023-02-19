package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	"data-platform-api-invoice-document-items-creates-subfunc/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"

	"golang.org/x/xerrors"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}
	data := pm

	res := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &res
}

func (psdc *SDC) ConvertToProcessType() *ProcessType {
	pm := &requests.ProcessType{}
	data := pm

	res := ProcessType{
		BulkProcess:       data.BulkProcess,
		IndividualProcess: data.IndividualProcess,
		ArraySpec:         data.ArraySpec,
		RangeSpec:         data.RangeSpec,
	}

	return &res
}

func (psdc *SDC) ConvertToReferenceType() *ReferenceType {
	pm := &requests.ReferenceType{}
	data := pm

	res := ReferenceType{
		OrderID:          data.OrderID,
		DeliveryDocument: data.DeliveryDocument,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderIDByArraySpecKey(length int) *OrderIDKey {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.BillFromParty = append(pm.BillFromParty, nil)
		pm.BillToParty = append(pm.BillToParty, nil)
	}

	data := pm
	res := OrderIDKey{
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		IsCancelled:                     data.IsCancelled,
		IsMarkedForDeletion:             data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderIDByArraySpec(rows *sql.Rows) ([]*OrderID, error) {
	defer rows.Close()
	res := make([]*OrderID, 0)

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		res = append(res, &OrderID{
			OrderID:                         data.OrderID,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		})
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpecKey() *OrderIDKey {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
	}

	data := pm
	res := OrderIDKey{
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		IsCancelled:                     data.IsCancelled,
		IsMarkedForDeletion:             data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderIDByRangeSpec(rows *sql.Rows) ([]*OrderID, error) {
	defer rows.Close()
	res := make([]*OrderID, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrderID{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderID{
			OrderID:                         data.OrderID,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
			IsCancelled:                     data.IsCancelled,
			IsMarkedForDeletion:             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderIDInIndividualProcessKey() *OrderIDKey {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
	}

	data := pm
	res := OrderIDKey{
		ReferenceDocument:               data.ReferenceDocument,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		IsCancelled:                     data.IsCancelled,
		IsMarkedForDeletion:             data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderIDInIndividualProcess(rows *sql.Rows) ([]*OrderID, error) {
	defer rows.Close()
	res := make([]*OrderID, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrderID{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderID{
			OrderID:                         data.OrderID,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
			IsCancelled:                     data.IsCancelled,
			IsMarkedForDeletion:             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderItemInBulkProcessKey() *OrderItemKey {
	pm := &requests.OrderItemKey{
		ItemCompleteDeliveryIsDefined: true,
		ItemDeliveryStatus:            "CL",
		ItemBillingStatus:             "CL",
		ItemBillingBlockStatus:        false,
		IsCancelled:                   getBoolPtr(false),
		IsMarkedForDeletion:           getBoolPtr(false),
	}

	data := pm
	res := OrderItemKey{
		OrderID:                       data.OrderID,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemBillingStatus:             data.ItemBillingStatus,
		ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
		IsCancelled:                   data.IsCancelled,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderItemInBulkProcess(rows *sql.Rows) ([]*OrderItem, error) {
	defer rows.Close()
	res := make([]*OrderItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrderItem{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemBillingStatus,
			&pm.ItemBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderItem{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemBillingStatus:             data.ItemBillingStatus,
			ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
			IsCancelled:                   data.IsCancelled,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrderItemInIndividualProcessKey() *OrderItemKey {
	pm := &requests.OrderItemKey{
		ItemCompleteDeliveryIsDefined: true,
		ItemDeliveryStatus:            "CL",
		ItemBillingStatus:             "CL",
		ItemBillingBlockStatus:        false,
		IsCancelled:                   getBoolPtr(false),
		IsMarkedForDeletion:           getBoolPtr(false),
	}

	data := pm
	res := OrderItemKey{
		OrderID:                       data.OrderID,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemBillingStatus:             data.ItemBillingStatus,
		ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
		IsCancelled:                   data.IsCancelled,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToOrderItemInIndividualProcess(rows *sql.Rows) ([]*OrderItem, error) {
	defer rows.Close()
	res := make([]*OrderItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrderItem{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemBillingStatus,
			&pm.ItemBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrderItem{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemBillingStatus:             data.ItemBillingStatus,
			ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
			IsCancelled:                   data.IsCancelled,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByArraySpecKey(length int) *DeliveryDocumentHeaderKey {
	pm := &requests.DeliveryDocumentHeaderKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
	}

	for i := 0; i < length; i++ {
		pm.BillFromParty = append(pm.BillFromParty, nil)
		pm.BillToParty = append(pm.BillToParty, nil)
	}

	data := pm
	res := DeliveryDocumentHeaderKey{
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		IsCancelled:                     data.IsCancelled,
		IsMarkedForDeletion:             data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToDeliveryDocumentByArraySpec(rows *sql.Rows) ([]*DeliveryDocumentHeader, error) {
	defer rows.Close()
	res := make([]*DeliveryDocumentHeader, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentHeader{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentHeader{
			DeliveryDocument:                data.DeliveryDocument,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
			IsCancelled:                     data.IsCancelled,
			IsMarkedForDeletion:             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentByRangeSpecKey() *DeliveryDocumentHeaderKey {
	pm := &requests.DeliveryDocumentHeaderKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
	}

	data := pm
	res := DeliveryDocumentHeaderKey{
		BillFromPartyFrom:               data.BillFromPartyFrom,
		BillFromPartyTo:                 data.BillFromPartyTo,
		BillToPartyFrom:                 data.BillToPartyFrom,
		BillToPartyTo:                   data.BillToPartyTo,
		BillFromParty:                   data.BillFromParty,
		BillToParty:                     data.BillToParty,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		IsCancelled:                     data.IsCancelled,
		IsMarkedForDeletion:             data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToDeliveryDocumentByRangeSpec(rows *sql.Rows) ([]*DeliveryDocumentHeader, error) {
	defer rows.Close()
	res := make([]*DeliveryDocumentHeader, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentHeader{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentHeader{
			DeliveryDocument:                data.DeliveryDocument,
			BillFromParty:                   data.BillFromParty,
			BillToParty:                     data.BillToParty,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
			IsCancelled:                     data.IsCancelled,
			IsMarkedForDeletion:             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentItemInBulkProcessKey() *DeliveryDocumentItemKey {
	pm := &requests.DeliveryDocumentItemKey{
		ItemCompleteDeliveryIsDefined: true,
		// ItemDeliveryStatus:            "CL",
		ItemBillingStatus:      "CL",
		ItemBillingBlockStatus: false,
		IsCancelled:            getBoolPtr(false),
		IsMarkedForDeletion:    getBoolPtr(false),
	}

	data := pm
	res := DeliveryDocumentItemKey{
		DeliveryDocument:              data.DeliveryDocument,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		// ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemBillingStatus:      data.ItemBillingStatus,
		ItemBillingBlockStatus: data.ItemBillingBlockStatus,
		IsCancelled:            data.IsCancelled,
		IsMarkedForDeletion:    data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToDeliveryDocumentItemInBulkProcess(rows *sql.Rows) ([]*DeliveryDocumentItem, error) {
	defer rows.Close()
	res := make([]*DeliveryDocumentItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentItem{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.ConfirmedDeliveryDate,
			&pm.ActualGoodsIssueDate,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemBillingStatus,
			&pm.ItemBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentItem{
			DeliveryDocument:              data.DeliveryDocument,
			DeliveryDocumentItem:          data.DeliveryDocumentItem,
			ConfirmedDeliveryDate:         data.ConfirmedDeliveryDate,
			ActualGoodsIssueDate:          data.ActualGoodsIssueDate,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemBillingStatus:             data.ItemBillingStatus,
			ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
			IsCancelled:                   data.IsCancelled,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentInIndividualProcessKey() *DeliveryDocumentHeaderKey {
	pm := &requests.DeliveryDocumentHeaderKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(true),
		HeaderDeliveryStatus:            "CL",
		HeaderBillingStatus:             "CL",
		HeaderBillingBlockStatus:        getBoolPtr(false),
		IsCancelled:                     getBoolPtr(false),
		IsMarkedForDeletion:             getBoolPtr(false),
	}

	data := pm
	res := DeliveryDocumentHeaderKey{
		ReferenceDocument:               data.ReferenceDocument,
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
		HeaderBillingStatus:             data.HeaderBillingStatus,
		HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
		IsCancelled:                     data.IsCancelled,
		IsMarkedForDeletion:             data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToDeliveryDocumentInIndividualProcess(rows *sql.Rows) ([]*DeliveryDocumentHeader, error) {
	defer rows.Close()
	res := make([]*DeliveryDocumentHeader, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentHeader{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.HeaderDeliveryStatus,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentHeader{
			DeliveryDocument:                data.DeliveryDocument,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			HeaderDeliveryStatus:            data.HeaderDeliveryStatus,
			HeaderBillingStatus:             data.HeaderBillingStatus,
			HeaderBillingBlockStatus:        data.HeaderBillingBlockStatus,
			IsCancelled:                     data.IsCancelled,
			IsMarkedForDeletion:             data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentItemInIndividualProcessKey() *DeliveryDocumentItemKey {
	pm := &requests.DeliveryDocumentItemKey{
		ItemCompleteDeliveryIsDefined: true,
		// ItemDeliveryStatus:            "CL",
		ItemBillingStatus:      "CL",
		ItemBillingBlockStatus: false,
		IsCancelled:            getBoolPtr(false),
		IsMarkedForDeletion:    getBoolPtr(false),
	}

	data := pm
	res := DeliveryDocumentItemKey{
		DeliveryDocument:              data.DeliveryDocument,
		ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
		// ItemDeliveryStatus:            data.ItemDeliveryStatus,
		ItemBillingStatus:      data.ItemBillingStatus,
		ItemBillingBlockStatus: data.ItemBillingBlockStatus,
		IsCancelled:            data.IsCancelled,
		IsMarkedForDeletion:    data.IsMarkedForDeletion,
	}

	return &res
}

func (psdc *SDC) ConvertToDeliveryDocumentItemInIndividualProcess(rows *sql.Rows) ([]*DeliveryDocumentItem, error) {
	defer rows.Close()
	res := make([]*DeliveryDocumentItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentItem{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.ConfirmedDeliveryDate,
			&pm.ActualGoodsIssueDate,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemBillingStatus,
			&pm.ItemBillingBlockStatus,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentItem{
			DeliveryDocument:              data.DeliveryDocument,
			DeliveryDocumentItem:          data.DeliveryDocumentItem,
			ConfirmedDeliveryDate:         data.ConfirmedDeliveryDate,
			ActualGoodsIssueDate:          data.ActualGoodsIssueDate,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemBillingStatus:             data.ItemBillingStatus,
			ItemBillingBlockStatus:        data.ItemBillingBlockStatus,
			IsCancelled:                   data.IsCancelled,
			IsMarkedForDeletion:           data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// Orders
func (psdc *SDC) ConvertToOrdersHeader(rows *sql.Rows) ([]*OrdersHeader, error) {
	defer rows.Close()
	res := make([]*OrdersHeader, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersHeader{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderType,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.TotalNetAmount,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.TransactionCurrency,
			&pm.PricingDate,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.IsExportImport,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersHeader{
			OrderID:                          data.OrderID,
			OrderType:                        data.OrderType,
			SupplyChainRelationshipID:        data.SupplyChainRelationshipID,
			SupplyChainRelationshipBillingID: data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID: data.SupplyChainRelationshipPaymentID,
			Buyer:                            data.Buyer,
			Seller:                           data.Seller,
			BillToParty:                      data.BillToParty,
			BillFromParty:                    data.BillFromParty,
			BillToCountry:                    data.BillToCountry,
			BillFromCountry:                  data.BillFromCountry,
			Payer:                            data.Payer,
			Payee:                            data.Payee,
			ContractType:                     data.ContractType,
			OrderValidityStartDate:           data.OrderValidityStartDate,
			OrderValidityEndDate:             data.OrderValidityEndDate,
			InvoicePeriodStartDate:           data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:             data.InvoicePeriodEndDate,
			TotalNetAmount:                   data.TotalNetAmount,
			TotalTaxAmount:                   data.TotalTaxAmount,
			TotalGrossAmount:                 data.TotalGrossAmount,
			TransactionCurrency:              data.TransactionCurrency,
			PricingDate:                      data.PricingDate,
			Incoterms:                        data.Incoterms,
			PaymentTerms:                     data.PaymentTerms,
			PaymentMethod:                    data.PaymentMethod,
			IsExportImport:                   data.IsExportImport,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersItem(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) ([]*OrdersItem, error) {

	defer rows.Close()
	res := make([]*OrdersItem, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersItem{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.OrderItemText,
			&pm.OrderItemTextByBuyer,
			&pm.OrderItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.PricingDate,
			&pm.DeliveryUnit,
			&pm.OrderQuantityInBaseUnit,
			&pm.OrderQuantityInDeliveryUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.Incoterms,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassification,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.Project,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.TaxCode,
			&pm.TaxRate,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersItem{
			OrderID:                                 data.OrderID,
			OrderItem:                               data.OrderItem,
			OrderItemCategory:                       data.OrderItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			OrderItemText:                           data.OrderItemText,
			OrderItemTextByBuyer:                    data.OrderItemTextByBuyer,
			OrderItemTextBySeller:                   data.OrderItemTextBySeller,
			Product:                                 data.Product,
			ProductStandardID:                       data.ProductStandardID,
			ProductGroup:                            data.ProductGroup,
			BaseUnit:                                data.BaseUnit,
			PricingDate:                             data.PricingDate,
			DeliveryUnit:                            data.DeliveryUnit,
			OrderQuantityInBaseUnit:                 data.OrderQuantityInBaseUnit,
			OrderQuantityInDeliveryUnit:             data.OrderQuantityInDeliveryUnit,
			NetAmount:                               data.NetAmount,
			TaxAmount:                               data.TaxAmount,
			GrossAmount:                             data.GrossAmount,
			Incoterms:                               data.Incoterms,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassification:                data.DefinedTaxClassification,
			PaymentTerms:                            data.PaymentTerms,
			PaymentMethod:                           data.PaymentMethod,
			Project:                                 data.Project,
			ReferenceDocument:                       data.ReferenceDocument,
			ReferenceDocumentItem:                   data.ReferenceDocumentItem,
			TaxCode:                                 data.TaxCode,
			TaxRate:                                 data.TaxRate,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToOrdersPartner(rows *sql.Rows) ([]*OrdersPartner, error) {
	defer rows.Close()
	res := make([]*OrdersPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.OrdersPartner{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &OrdersPartner{
			OrderID:                 data.OrderID,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_partner_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToItemPricingElement(rows *sql.Rows) ([]*ItemPricingElement, error) {
	defer rows.Close()
	res := make([]*ItemPricingElement, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.ItemPricingElement{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.PricingProcedureCounter,
			&pm.ConditionRecord,
			&pm.ConditionSequentialNumber,
			&pm.ConditionType,
			&pm.PricingDate,
			&pm.ConditionRateValue,
			&pm.ConditionCurrency,
			&pm.ConditionQuantity,
			&pm.ConditionQuantityUnit,
			&pm.ConditionAmount,
			&pm.ConditionIsManuallyChanged,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &ItemPricingElement{
			OrderID:                    data.OrderID,
			OrderItem:                  data.OrderItem,
			PricingProcedureCounter:    data.PricingProcedureCounter,
			ConditionRecord:            data.ConditionRecord,
			ConditionSequentialNumber:  data.ConditionSequentialNumber,
			ConditionType:              data.ConditionType,
			PricingDate:                data.PricingDate,
			ConditionRateValue:         data.ConditionRateValue,
			ConditionCurrency:          data.ConditionCurrency,
			ConditionQuantity:          data.ConditionQuantity,
			ConditionQuantityUnit:      data.ConditionQuantityUnit,
			ConditionAmount:            data.ConditionAmount,
			ConditionIsManuallyChanged: data.ConditionIsManuallyChanged,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_orders_item_pricing_element_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// DeliveryDocument
func (psdc *SDC) ConvertToDeliveryDocumentHeaderData(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) ([]*DeliveryDocumentHeaderData, error) {

	defer rows.Close()
	res := make([]*DeliveryDocumentHeaderData, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentHeaderData{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.IsExportImport,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.GoodsIssueOrReceiptSlipNumber,
			&pm.Incoterms,
			&pm.TransactionCurrency,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentHeaderData{
			DeliveryDocument:                       data.DeliveryDocument,
			SupplyChainRelationshipID:              data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:      data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID: data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipBillingID:       data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID:       data.SupplyChainRelationshipPaymentID,
			Buyer:                                  data.Buyer,
			Seller:                                 data.Seller,
			DeliverToParty:                         data.DeliverToParty,
			DeliverFromParty:                       data.DeliverFromParty,
			DeliverToPlant:                         data.DeliverToPlant,
			DeliverFromPlant:                       data.DeliverFromPlant,
			BillToParty:                            data.BillToParty,
			BillFromParty:                          data.BillFromParty,
			BillToCountry:                          data.BillToCountry,
			BillFromCountry:                        data.BillFromCountry,
			Payer:                                  data.Payer,
			Payee:                                  data.Payee,
			IsExportImport:                         data.IsExportImport,
			OrderID:                                data.OrderID,
			OrderItem:                              data.OrderItem,
			ContractType:                           data.ContractType,
			OrderValidityStartDate:                 data.OrderValidityStartDate,
			OrderValidityEndDate:                   data.OrderValidityEndDate,
			GoodsIssueOrReceiptSlipNumber:          data.GoodsIssueOrReceiptSlipNumber,
			Incoterms:                              data.Incoterms,
			TransactionCurrency:                    data.TransactionCurrency,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_header_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentItemData(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) ([]*DeliveryDocumentItemData, error) {

	defer rows.Close()
	res := make([]*DeliveryDocumentItemData, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentItemData{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.DeliveryDocumentItem,
			&pm.DeliveryDocumentItemCategory,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipBillingID,
			&pm.SupplyChainRelationshipPaymentID,
			&pm.Buyer,
			&pm.Seller,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromPlant,
			&pm.BillToParty,
			&pm.BillFromParty,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.DeliverToPlantStorageLocation,
			&pm.DeliverFromPlantStorageLocation,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.ProductionPlantStorageLocation,
			&pm.DeliveryDocumentItemText,
			&pm.DeliveryDocumentItemTextByBuyer,
			&pm.DeliveryDocumentItemTextBySeller,
			&pm.Product,
			&pm.ProductStandardID,
			&pm.ProductGroup,
			&pm.BaseUnit,
			&pm.DeliveryUnit,
			&pm.ActualGoodsIssueDate,
			&pm.ActualGoodsIssueTime,
			&pm.ActualGoodsReceiptDate,
			&pm.ActualGoodsReceiptTime,
			&pm.ActualGoodsIssueQuantity,
			&pm.ActualGoodsIssueQtyInBaseUnit,
			&pm.ActualGoodsReceiptQuantity,
			&pm.ActualGoodsReceiptQtyInBaseUnit,
			&pm.ItemGrossWeight,
			&pm.ItemNetWeight,
			&pm.ItemWeightUnit,
			&pm.NetAmount,
			&pm.TaxAmount,
			&pm.GrossAmount,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.OrderType,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.InvoicePeriodStartDate,
			&pm.InvoicePeriodEndDate,
			&pm.Project,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.TransactionTaxClassification,
			&pm.ProductTaxClassificationBillToCountry,
			&pm.ProductTaxClassificationBillFromCountry,
			&pm.DefinedTaxClassifications,
			&pm.TaxCode,
			&pm.TaxRate,
			&pm.CountryOfOrigin,
			&pm.CountryOfOriginLanguage,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentItemData{
			DeliveryDocument:                        data.DeliveryDocument,
			DeliveryDocumentItem:                    data.DeliveryDocumentItem,
			DeliveryDocumentItemCategory:            data.DeliveryDocumentItemCategory,
			SupplyChainRelationshipID:               data.SupplyChainRelationshipID,
			SupplyChainRelationshipDeliveryID:       data.SupplyChainRelationshipDeliveryID,
			SupplyChainRelationshipDeliveryPlantID:  data.SupplyChainRelationshipDeliveryPlantID,
			SupplyChainRelationshipBillingID:        data.SupplyChainRelationshipBillingID,
			SupplyChainRelationshipPaymentID:        data.SupplyChainRelationshipPaymentID,
			Buyer:                                   data.Buyer,
			Seller:                                  data.Seller,
			DeliverToParty:                          data.DeliverToParty,
			DeliverFromParty:                        data.DeliverFromParty,
			DeliverToPlant:                          data.DeliverToPlant,
			DeliverFromPlant:                        data.DeliverFromPlant,
			BillToParty:                             data.BillToParty,
			BillFromParty:                           data.BillFromParty,
			BillToCountry:                           data.BillToCountry,
			BillFromCountry:                         data.BillFromCountry,
			Payer:                                   data.Payer,
			Payee:                                   data.Payee,
			DeliverToPlantStorageLocation:           data.DeliverToPlantStorageLocation,
			DeliverFromPlantStorageLocation:         data.DeliverFromPlantStorageLocation,
			ProductionPlantBusinessPartner:          data.ProductionPlantBusinessPartner,
			ProductionPlant:                         data.ProductionPlant,
			ProductionPlantStorageLocation:          data.ProductionPlantStorageLocation,
			DeliveryDocumentItemText:                data.DeliveryDocumentItemText,
			DeliveryDocumentItemTextByBuyer:         data.DeliveryDocumentItemTextByBuyer,
			DeliveryDocumentItemTextBySeller:        data.DeliveryDocumentItemTextBySeller,
			Product:                                 data.Product,
			ProductStandardID:                       data.ProductStandardID,
			ProductGroup:                            data.ProductGroup,
			BaseUnit:                                data.BaseUnit,
			DeliveryUnit:                            data.DeliveryUnit,
			ActualGoodsIssueDate:                    data.ActualGoodsIssueDate,
			ActualGoodsIssueTime:                    data.ActualGoodsIssueTime,
			ActualGoodsReceiptDate:                  data.ActualGoodsReceiptDate,
			ActualGoodsReceiptTime:                  data.ActualGoodsReceiptTime,
			ActualGoodsIssueQuantity:                data.ActualGoodsIssueQuantity,
			ActualGoodsIssueQtyInBaseUnit:           data.ActualGoodsIssueQtyInBaseUnit,
			ActualGoodsReceiptQuantity:              data.ActualGoodsReceiptQuantity,
			ActualGoodsReceiptQtyInBaseUnit:         data.ActualGoodsReceiptQtyInBaseUnit,
			ItemGrossWeight:                         data.ItemGrossWeight,
			ItemNetWeight:                           data.ItemNetWeight,
			ItemWeightUnit:                          data.ItemWeightUnit,
			NetAmount:                               data.NetAmount,
			TaxAmount:                               data.TaxAmount,
			GrossAmount:                             data.GrossAmount,
			OrderID:                                 data.OrderID,
			OrderItem:                               data.OrderItem,
			OrderType:                               data.OrderType,
			ContractType:                            data.ContractType,
			OrderValidityStartDate:                  data.OrderValidityStartDate,
			OrderValidityEndDate:                    data.OrderValidityEndDate,
			PaymentTerms:                            data.PaymentTerms,
			PaymentMethod:                           data.PaymentMethod,
			InvoicePeriodStartDate:                  data.InvoicePeriodStartDate,
			InvoicePeriodEndDate:                    data.InvoicePeriodEndDate,
			Project:                                 data.Project,
			ReferenceDocument:                       data.ReferenceDocument,
			ReferenceDocumentItem:                   data.ReferenceDocumentItem,
			TransactionTaxClassification:            data.TransactionTaxClassification,
			ProductTaxClassificationBillToCountry:   data.ProductTaxClassificationBillToCountry,
			ProductTaxClassificationBillFromCountry: data.ProductTaxClassificationBillFromCountry,
			DefinedTaxClassifications:               data.DefinedTaxClassifications,
			TaxCode:                                 data.TaxCode,
			TaxRate:                                 data.TaxRate,
			CountryOfOrigin:                         data.CountryOfOrigin,
			CountryOfOriginLanguage:                 data.CountryOfOriginLanguage,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_item_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentPartner(rows *sql.Rows) ([]*DeliveryDocumentPartner, error) {
	defer rows.Close()
	res := make([]*DeliveryDocumentPartner, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.DeliveryDocumentPartner{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &DeliveryDocumentPartner{
			DeliveryDocument:        data.DeliveryDocument,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_delivery_document_partner_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

// Header
func (psdc *SDC) ConvertToInvoiceDocumentHeaderKey() *CalculateInvoiceDocumentKey {
	pm := &requests.CalculateInvoiceDocumentKey{
		FieldNameWithNumberRange: "InvoiceDocument",
	}

	data := pm
	res := CalculateInvoiceDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &res
}

func (psdc *SDC) ConvertToInvoiceDocumentHeaderQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*CalculateInvoiceDocumentQueryGets, error) {
	defer rows.Close()
	var res *CalculateInvoiceDocumentQueryGets

	i := 0
	for rows.Next() {
		i++
		pm := &requests.CalculateInvoiceDocumentQueryGets{}

		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.InvoiceDocumentLatestNumber,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = &CalculateInvoiceDocumentQueryGets{
			ServiceLabel:                data.ServiceLabel,
			FieldNameWithNumberRange:    data.FieldNameWithNumberRange,
			InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
		}

	}
	if i == 0 {
		return nil, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocument(invoiceDocumentLatestNumber *int, invoiceDocument, orderID, deliveryDocument, deliveryDocumentItem, billFromParty, billToParty int) *CalculateInvoiceDocument {
	pm := &requests.CalculateInvoiceDocument{}

	pm.InvoiceDocumentLatestNumber = invoiceDocumentLatestNumber
	pm.InvoiceDocument = invoiceDocument
	pm.OrderID = orderID
	pm.DeliveryDocument = deliveryDocument
	pm.DeliveryDocumentItem = deliveryDocumentItem
	pm.BillFromParty = billFromParty
	pm.BillToParty = billToParty

	data := pm
	res := CalculateInvoiceDocument{
		InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
		InvoiceDocument:             data.InvoiceDocument,
		OrderID:                     data.OrderID,
		OrderItem:                   data.OrderItem,
		DeliveryDocument:            data.DeliveryDocument,
		DeliveryDocumentItem:        data.DeliveryDocumentItem,
		BillFromParty:               data.BillFromParty,
		BillToParty:                 data.BillToParty,
	}

	return &res
}

// Item
func (psdc *SDC) ConvertToInvoiceDocumentItem(sdc *api_input_reader.SDC) []*InvoiceDocumentItem {
	res := make([]*InvoiceDocumentItem, 0)

	processType := psdc.ProcessType
	referenceType := psdc.ReferenceType
	if referenceType.OrderID {
		if processType.BulkProcess {
			for i, orderItem := range psdc.OrderItem {
				pm := &requests.InvoiceDocumentItem{}

				pm.OrderID = orderItem.OrderID
				pm.OrderItem = orderItem.OrderItem
				pm.InvoiceDocumentItemNumber = i + 1

				data := pm
				res = append(res, &InvoiceDocumentItem{
					OrderID:                   data.OrderID,
					OrderItem:                 data.OrderItem,
					InvoiceDocumentItemNumber: data.InvoiceDocumentItemNumber,
				})
			}
		} else if processType.IndividualProcess {
			for i, item := range sdc.Header.Item {
				pm := &requests.InvoiceDocumentItem{}

				pm.OrderID = *item.OrderID
				pm.OrderItem = *item.OrderItem
				pm.InvoiceDocumentItemNumber = i + 1

				data := pm
				res = append(res, &InvoiceDocumentItem{
					OrderID:                   data.OrderID,
					OrderItem:                 data.OrderItem,
					InvoiceDocumentItemNumber: data.InvoiceDocumentItemNumber,
				})
			}
		}
	} else if referenceType.DeliveryDocument {
		if processType.BulkProcess {
			for i, deliveryDocumentItem := range psdc.DeliveryDocumentItem {
				pm := &requests.InvoiceDocumentItem{}

				pm.DeliveryDocument = deliveryDocumentItem.DeliveryDocument
				pm.DeliveryDocumentItem = deliveryDocumentItem.DeliveryDocumentItem
				pm.InvoiceDocumentItemNumber = i + 1

				data := pm
				res = append(res, &InvoiceDocumentItem{
					DeliveryDocument:          data.DeliveryDocument,
					DeliveryDocumentItem:      data.DeliveryDocumentItem,
					InvoiceDocumentItemNumber: data.InvoiceDocumentItemNumber,
				})
			}
		} else if processType.IndividualProcess {
			for i, item := range sdc.Header.Item {
				pm := &requests.InvoiceDocumentItem{}

				pm.DeliveryDocument = *item.DeliveryDocument
				pm.DeliveryDocumentItem = *item.DeliveryDocumentItem
				pm.InvoiceDocumentItemNumber = i + 1

				data := pm
				res = append(res, &InvoiceDocumentItem{
					DeliveryDocument:          data.DeliveryDocument,
					DeliveryDocumentItem:      data.DeliveryDocumentItem,
					InvoiceDocumentItemNumber: data.InvoiceDocumentItemNumber,
				})
			}
		}
	}

	return res
}

// Address
func (psdc *SDC) ConvertToOrdersAddress(rows *sql.Rows) ([]*Address, error) {
	defer rows.Close()
	res := make([]*Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.OrderID,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &Address{
			OrderID:     data.OrderID,
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_orders_address_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentAddress(rows *sql.Rows) ([]*Address, error) {
	defer rows.Close()
	res := make([]*Address, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Address{}

		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.AddressID,
			&pm.PostalCode,
			&pm.LocalRegion,
			&pm.Country,
			&pm.District,
			&pm.StreetName,
			&pm.CityName,
			&pm.Building,
			&pm.Floor,
			&pm.Room,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		res = append(res, &Address{
			DeliveryDocument: data.DeliveryDocument,
			AddressID:        data.AddressID,
			PostalCode:       data.PostalCode,
			LocalRegion:      data.LocalRegion,
			Country:          data.Country,
			District:         data.District,
			StreetName:       data.StreetName,
			CityName:         data.CityName,
			Building:         data.Building,
			Floor:            data.Floor,
			Room:             data.Room,
		})
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_delivery_document_address_data'テーブルに対象のレコードが存在しません。")
	}

	return res, nil
}

func (psdc *SDC) ConvertToAddressMaster(sdc *api_input_reader.SDC, idx, addressID int) *AddressMaster {
	pm := &requests.AddressMaster{
		// ValidityStartDate: *sdc.Header.OrderValidityStartDate,
		// ValidityEndDate:   *sdc.Header.OrderValidityEndDate,
		PostalCode:  *sdc.Header.Address[idx].PostalCode,
		LocalRegion: *sdc.Header.Address[idx].LocalRegion,
		Country:     *sdc.Header.Address[idx].Country,
		District:    sdc.Header.Address[idx].District,
		StreetName:  *sdc.Header.Address[idx].StreetName,
		CityName:    *sdc.Header.Address[idx].CityName,
		Building:    sdc.Header.Address[idx].Building,
		Floor:       sdc.Header.Address[idx].Floor,
		Room:        sdc.Header.Address[idx].Room,
	}

	pm.AddressID = addressID

	data := pm
	res := &AddressMaster{
		AddressID:         data.AddressID,
		ValidityEndDate:   data.ValidityEndDate,
		ValidityStartDate: data.ValidityStartDate,
		PostalCode:        data.PostalCode,
		LocalRegion:       data.LocalRegion,
		Country:           data.Country,
		District:          data.District,
		StreetName:        data.StreetName,
		CityName:          data.CityName,
		Building:          data.Building,
		Floor:             data.Floor,
		Room:              data.Room,
	}

	return res
}

func (psdc *SDC) ConvertToAddressFromInput() []*Address {
	res := make([]*Address, 0)

	for _, v := range psdc.AddressMaster {
		pm := &requests.Address{}

		pm.AddressID = v.AddressID
		pm.PostalCode = &v.PostalCode
		pm.LocalRegion = &v.LocalRegion
		pm.Country = &v.Country
		pm.District = v.District
		pm.StreetName = &v.StreetName
		pm.CityName = &v.CityName
		pm.Building = v.Building
		pm.Floor = v.Floor
		pm.Room = v.Room

		data := pm
		res = append(res, &Address{
			AddressID:   data.AddressID,
			PostalCode:  data.PostalCode,
			LocalRegion: data.LocalRegion,
			Country:     data.Country,
			District:    data.District,
			StreetName:  data.StreetName,
			CityName:    data.CityName,
			Building:    data.Building,
			Floor:       data.Floor,
			Room:        data.Room,
		})

	}

	return res
}

func (psdc *SDC) ConvertToCalculateAddressIDKey() *CalculateAddressIDKey {
	pm := &requests.CalculateAddressIDKey{
		ServiceLabel:             "ADDRESS_ID",
		FieldNameWithNumberRange: "AddressID",
	}

	data := pm
	res := CalculateAddressIDKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &res
}

func (psdc *SDC) ConvertToCalculateAddressIDQueryGets(rows *sql.Rows) (*CalculateAddressIDQueryGets, error) {
	defer rows.Close()
	pm := &requests.CalculateAddressIDQueryGets{}

	i := 0
	for rows.Next() {
		i++
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.LatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	if i == 0 {
		return nil, xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
	}

	data := pm
	res := CalculateAddressIDQueryGets{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
		LatestNumber:             data.LatestNumber,
	}

	return &res, nil
}

func (psdc *SDC) ConvertToCalculateAddressID(addressIDLatestNumber *int, addressID int) *CalculateAddressID {
	pm := &requests.CalculateAddressID{}

	pm.AddressIDLatestNumber = addressIDLatestNumber
	pm.AddressID = addressID

	data := pm

	res := CalculateAddressID{
		AddressIDLatestNumber: data.AddressIDLatestNumber,
		AddressID:             data.AddressID,
	}

	return &res
}

// 日付等の処理
func (psdc *SDC) ConvertToCreationDateItem(systemDate string) *CreationDate {
	pm := &requests.CreationDate{}

	pm.CreationDate = systemDate

	data := pm
	res := CreationDate{
		CreationDate: data.CreationDate,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeDateItem(systemDate string) *LastChangeDate {
	pm := &requests.LastChangeDate{}

	pm.LastChangeDate = systemDate

	data := pm
	res := LastChangeDate{
		LastChangeDate: data.LastChangeDate,
	}

	return &res
}

func (psdc *SDC) ConvertToCreationTimeItem(systemTime string) *CreationTime {
	pm := &requests.CreationTime{}

	pm.CreationTime = systemTime

	data := pm
	res := CreationTime{
		CreationTime: data.CreationTime,
	}

	return &res
}

func (psdc *SDC) ConvertToLastChangeTimeItem(systemTime string) *LastChangeTime {
	pm := &requests.LastChangeTime{}

	pm.LastChangeTime = systemTime

	data := pm
	res := LastChangeTime{
		LastChangeTime: data.LastChangeTime,
	}

	return &res
}

func getBoolPtr(b bool) *bool {
	return &b
}
