package subfunction

import (
	"context"
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"data-platform-api-invoice-document-items-creates-subfunc/config"
	"strings"

	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SubFunction struct {
	ctx  context.Context
	db   *database.Mysql
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	l    *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, conf *config.Conf, rmq *rabbitmq.RabbitmqClient, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx:  ctx,
		db:   db,
		conf: conf,
		rmq:  rmq,
		l:    l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) ProcessType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.ProcessType {
	processType := psdc.ConvertToProcessType()

	processType.BulkProcess = true
	// processType.IndividualProcess = true

	if processType.BulkProcess {
		// processType.ArraySpec = true
		processType.RangeSpec = true
	}

	return processType
}

func (f *SubFunction) ReferenceType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.ReferenceType {
	referenceType := psdc.ConvertToReferenceType()

	referenceType.OrderID = true
	// referenceType.DeliveryDocument = true

	return referenceType
}

func (f *SubFunction) OrderIDInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderID, error) {
	data := make([]*api_processing_data_formatter.OrderID, 0)
	var err error

	processType := psdc.ProcessType

	if processType.ArraySpec {
		data, err = f.OrderIDByArraySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.OrderIDByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (f *SubFunction) OrderIDByArraySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderID, error) {
	args := make([]interface{}, 0)

	billFromParty := sdc.InputParameters.BillFromParty
	billToParty := sdc.InputParameters.BillToParty

	if len(*billFromParty) != len(*billToParty) {
		return nil, nil
	}

	dataKey := psdc.ConvertToOrderIDByArraySpecKey(len(*billFromParty))

	for i := range *billFromParty {
		dataKey.BillFromParty[i] = (*billFromParty)[i]
		dataKey.BillToParty[i] = (*billToParty)[i]
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.BillFromParty)-1) + "(?,?)"
	for i := range dataKey.BillFromParty {
		args = append(args, dataKey.BillFromParty[i], dataKey.BillToParty[i])
	}

	args = append(
		args,
		dataKey.HeaderCompleteDeliveryIsDefined,
		dataKey.HeaderDeliveryStatus,
		dataKey.HeaderBillingBlockStatus,
		dataKey.HeaderBillingStatus,
		dataKey.IsCancelled,
		dataKey.IsMarkedForDeletion,
	)

	var count *int
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderIDの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderIDByArraySpec(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderID, error) {
	dataKey := psdc.ConvertToOrderIDByRangeSpecKey()

	dataKey.BillFromPartyFrom = sdc.InputParameters.BillFromPartyFrom
	dataKey.BillFromPartyTo = sdc.InputParameters.BillFromPartyTo
	dataKey.BillToPartyFrom = sdc.InputParameters.BillToPartyFrom
	dataKey.BillToPartyTo = sdc.InputParameters.BillToPartyTo

	count := new(int)
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderIDの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus,IsCancelled, IsMarkedForDeletion, HeaderBillingBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderIDByRangeSpec(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderIDInIndividualProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderID, error) {
	dataKey := psdc.ConvertToOrderIDInIndividualProcessKey()

	dataKey.ReferenceDocument = sdc.InputParameters.ReferenceDocument

	rows, err := f.db.Query(
		`SELECT OrderID, HeaderCompleteDeliveryIsDefined, HeaderBillingStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (OrderID, HeaderCompleteDeliveryIsDefined, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.ReferenceDocument, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderIDInIndividualProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrderItemInBulkProcessKey()

	orderID := psdc.OrderID

	for i := range orderID {
		dataKey.OrderID = append(dataKey.OrderID, (orderID)[i].OrderID)
	}

	repeat := strings.Repeat("?,", len(dataKey.OrderID)-1) + "?"
	for _, v := range dataKey.OrderID {
		args = append(args, v)
	}

	args = append(args, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryStatus, dataKey.ItemBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.ItemBillingStatus)

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBillingStatus, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE OrderID IN ( `+repeat+` )
		AND (ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND ItemBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItemInBulkProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemInIndividualProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrderItemInIndividualProcessKey()

	orderID := psdc.OrderID

	for i := range orderID {
		dataKey.OrderID = append(dataKey.OrderID, (orderID)[i].OrderID)
	}

	repeat := strings.Repeat("?,", len(dataKey.OrderID)-1) + "?"
	for _, v := range dataKey.OrderID {
		args = append(args, v)
	}

	args = append(args, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryStatus, dataKey.ItemBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.ItemBillingStatus)

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBillingStatus, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE OrderID IN ( `+repeat+` )
		AND (ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND ItemBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItemInIndividualProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentHeader, error) {
	data := make([]*api_processing_data_formatter.DeliveryDocumentHeader, 0)
	var err error

	processType := psdc.ProcessType

	if processType.ArraySpec {
		data, err = f.DeliveryDocumentByArraySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.DeliveryDocumentByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (f *SubFunction) DeliveryDocumentByArraySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentHeader, error) {
	args := make([]interface{}, 0)

	billFromParty := sdc.InputParameters.BillFromParty
	billToParty := sdc.InputParameters.BillToParty

	if len(*billFromParty) != len(*billToParty) {
		return nil, nil
	}

	dataKey := psdc.ConvertToDeliveryDocumentByArraySpecKey(len(*billFromParty))

	for i := range *billFromParty {
		dataKey.BillFromParty[i] = (*billFromParty)[i]
		dataKey.BillToParty[i] = (*billToParty)[i]
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.BillFromParty)-1) + "(?,?)"
	for i := range dataKey.BillFromParty {
		args = append(args, dataKey.BillFromParty[i], dataKey.BillToParty[i])
	}

	args = append(
		args,
		dataKey.HeaderCompleteDeliveryIsDefined,
		dataKey.HeaderDeliveryStatus,
		dataKey.HeaderBillingBlockStatus,
		dataKey.IsCancelled,
		dataKey.IsMarkedForDeletion,
		dataKey.HeaderBillingStatus,
	)

	count := new(int)
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (BillFromParty, BillToParty) IN ( `+repeat+` )
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, args...,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("DeliveryDocumentの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND  (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentByArraySpec(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentHeader, error) {
	dataKey := psdc.ConvertToDeliveryDocumentByRangeSpecKey()

	dataKey.BillFromPartyFrom = sdc.InputParameters.BillFromPartyFrom
	dataKey.BillFromPartyTo = sdc.InputParameters.BillFromPartyTo
	dataKey.BillToPartyFrom = sdc.InputParameters.BillToPartyFrom
	dataKey.BillToPartyTo = sdc.InputParameters.BillToPartyTo

	count := new(int)
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("DeliveryDocumentの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, BillFromParty, BillToParty, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE BillFromParty BETWEEN ? AND ?
		AND BillToParty BETWEEN ? AND ?
		AND  (HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.BillFromPartyFrom, dataKey.BillFromPartyTo, dataKey.BillToPartyFrom, dataKey.BillToPartyTo, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentByRangeSpec(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentItemInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToDeliveryDocumentItemInBulkProcessKey()

	dataKey.ConfirmedDeliveryDateFrom = sdc.InputParameters.ConfirmedDeliveryDateFrom
	dataKey.ConfirmedDeliveryDateTo = sdc.InputParameters.ConfirmedDeliveryDateTo
	dataKey.ActualGoodsIssueDateFrom = sdc.InputParameters.ActualGoodsIssueDateFrom
	dataKey.ActualGoodsIssueDateTo = sdc.InputParameters.ActualGoodsIssueDateTo

	deliveryDocumentItem := psdc.DeliveryDocumentHeader

	for i := range deliveryDocumentItem {
		dataKey.DeliveryDocument = append(dataKey.DeliveryDocument, deliveryDocumentItem[i].DeliveryDocument)
	}

	repeat := strings.Repeat("?,", len(dataKey.DeliveryDocument)-1) + "?"
	for _, v := range dataKey.DeliveryDocument {
		args = append(args, v)
	}

	args = append(args, dataKey.ConfirmedDeliveryDateFrom, dataKey.ConfirmedDeliveryDateTo, dataKey.ActualGoodsIssueDateFrom, dataKey.ActualGoodsIssueDateTo, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.ItemBillingStatus)

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, DeliveryDocumentItem, ConfirmedDeliveryDate, ActualGoodsIssueDate, ItemCompleteDeliveryIsDefined, ItemBillingStatus, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		WHERE DeliveryDocument IN ( `+repeat+` )
		AND ConfirmedDeliveryDate BETWEEN ? AND ?
		AND ActualGoodsIssueDate BETWEEN ? AND ?
		AND (ItemCompleteDeliveryIsDefined, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?)
		AND ItemBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentItemInBulkProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentInIndividualProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentHeader, error) {
	dataKey := psdc.ConvertToDeliveryDocumentInIndividualProcessKey()

	dataKey.ReferenceDocument = sdc.InputParameters.ReferenceDocument

	count := new(int)
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (DeliveryDocument, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.ReferenceDocument, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 {
		return nil, xerrors.Errorf("OrderIDの検索結果がゼロ件です。")
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus,  HeaderBillingStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE (DeliveryDocument, HeaderCompleteDeliveryIsDefined, HeaderDeliveryStatus, HeaderBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?, ?, ?)
		AND HeaderBillingStatus <> ?;`, dataKey.ReferenceDocument, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.HeaderDeliveryStatus, dataKey.HeaderBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.HeaderBillingStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentInIndividualProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentItemInIndividualProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.DeliveryDocumentItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToDeliveryDocumentItemInIndividualProcessKey()

	dataKey.ConfirmedDeliveryDateFrom = sdc.InputParameters.ConfirmedDeliveryDateFrom
	dataKey.ConfirmedDeliveryDateTo = sdc.InputParameters.ConfirmedDeliveryDateTo
	dataKey.ActualGoodsIssueDateFrom = sdc.InputParameters.ActualGoodsIssueDateFrom
	dataKey.ActualGoodsIssueDateTo = sdc.InputParameters.ActualGoodsIssueDateTo

	deliveryDocumentItem := psdc.DeliveryDocumentHeader

	for i := range deliveryDocumentItem {
		dataKey.DeliveryDocument = append(dataKey.DeliveryDocument, deliveryDocumentItem[i].DeliveryDocument)
	}

	repeat := strings.Repeat("?,", len(dataKey.DeliveryDocument)-1) + "?"
	for _, v := range dataKey.DeliveryDocument {
		args = append(args, v)
	}

	args = append(args, dataKey.ConfirmedDeliveryDateFrom, dataKey.ConfirmedDeliveryDateTo, dataKey.ActualGoodsIssueDateFrom, dataKey.ActualGoodsIssueDateTo, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemBillingBlockStatus, dataKey.IsCancelled, dataKey.IsMarkedForDeletion, dataKey.ItemBillingStatus)

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, DeliveryDocumentItem, ConfirmedDeliveryDate, ActualGoodsIssueDate, ItemCompleteDeliveryIsDefined, ItemBillingStatus, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_item_data
		WHERE DeliveryDocument IN ( `+repeat+` )
		AND ConfirmedDeliveryDate BETWEEN ? AND ?
		AND ActualGoodsIssueDate BETWEEN ? AND ?
		AND (ItemCompleteDeliveryIsDefined, ItemBillingBlockStatus, IsCancelled, IsMarkedForDeletion) = (?, ?, ?, ?)
		AND ItemBillingStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentItemInIndividualProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error

	psdc.MetaData = f.MetaData(sdc, psdc)
	psdc.ProcessType = f.ProcessType(sdc, psdc)
	psdc.ReferenceType = f.ReferenceType(sdc, psdc)

	referenceType := psdc.ReferenceType

	if referenceType.OrderID {
		err = f.OrdersReferenceProcess(sdc, psdc, osdc)
		if err != nil {
			return err
		}
	} else if referenceType.DeliveryDocument {
		err = f.DeliveryDocumentReferenceProcess(sdc, psdc, osdc)
		if err != nil {
			return err
		}
	}

	f.l.Info(psdc)
	err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}

func (f *SubFunction) OrdersReferenceProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	psdc.MetaData = f.MetaData(sdc, psdc)
	psdc.ProcessType = f.ProcessType(sdc, psdc)

	processType := psdc.ProcessType

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if processType.BulkProcess {
			// I-1-1. OrderIDの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
			psdc.OrderID, e = f.OrderIDInBulkProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// I-1-2. OrderItemの絞り込み  //I-1-1
			psdc.OrderItem, e = f.OrderItemInBulkProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		} else if processType.IndividualProcess {
			// II-1-1. OrderIDが未請求対象であることの確認
			psdc.OrderID, e = f.OrderIDInIndividualProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// II-1-2. OrderItemの絞り込み  //II-1-1
			psdc.OrderItem, e = f.OrderItemInIndividualProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 1-1.オーダー参照レコード・値の取得（オーダーヘッダ）  //I-1-1
			psdc.OrdersHeader, e = f.OrdersHeader(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 1-40. オーダー参照レコード・値の取得（オーダーパートナ）  //1-1
				psdc.OrdersPartner, e = f.OrdersPartner(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 6-1-1. Orders Address からの住所データの取得  //1-1
				psdc.Address, e = f.OrdersAddress(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 6-2. AddressIDの登録(ユーザーが任意の住所を入力ファイルで指定した場合)
				psdc.Address, e = f.AddressFromInput(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)

			// InvoiceDocumentItem  //(I-1-2またはII-1-2)
			psdc.InvoiceDocumentItem = f.InvoiceDocumentItem(sdc, psdc)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				//3-1. InvoiceDocumentHeader  //I-1-1
				psdc.CalculateInvoiceDocument, e = f.CalculateInvoiceDocument(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 1-2. オーダー参照レコード・値の取得（オーダー明細）  //II-1-1
			psdc.OrdersItem, e = f.OrdersItem(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 4-1. オーダー参照の場合の価格決定要素データの取得
			psdc.ItemPricingElement, e = f.OrdersItemPricingElement(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}(wg)

	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 99-1-2. CreationDate(Item)
		psdc.CreationDateItem = f.CreationDateItem(sdc, psdc)

		// 99-2-2. LastChangeDate(Item)
		psdc.LastChangeDateItem = f.LastChangeDateItem(sdc, psdc)

		// 99-3-2. CrationTime(Item)
		psdc.CreationTimeItem = f.CreationTimeItem(sdc, psdc)

		// 99-4-2. LastChangeTimeItem(Item)
		psdc.LastChangeTimeItem = f.LastChangeTimeItem(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (f *SubFunction) DeliveryDocumentReferenceProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	psdc.MetaData = f.MetaData(sdc, psdc)
	psdc.ProcessType = f.ProcessType(sdc, psdc)

	processType := psdc.ProcessType

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if processType.BulkProcess {
			// I-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
			psdc.DeliveryDocumentHeader, e = f.DeliveryDocumentInBulkProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// I-2-2. Delivery Document Itemの絞り込み  //I-2-1
			psdc.DeliveryDocumentItem, e = f.DeliveryDocumentItemInBulkProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		} else if processType.IndividualProcess {
			// II-2-1. Delivery Document Headerの絞り込み、および、入力パラメータによる請求元と請求先の絞り込み
			psdc.DeliveryDocumentHeader, e = f.DeliveryDocumentInIndividualProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// II-2-2. Delivery Document Itemの絞り込み  //II-2-1
			psdc.DeliveryDocumentItem, e = f.DeliveryDocumentItemInIndividualProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 2-1. 入出荷伝票参照レコード・値の取得（入出荷伝票ヘッダ）  //I-2-2
			psdc.DeliveryDocumentHeaderData, e = f.DeliveryDocumentHeaderData(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-2. 入出荷伝票参照レコード・値の取得（入出荷伝票明細）  //I-2-2
			psdc.DeliveryDocumentItemData, e = f.DeliveryDocumentItemData(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 4-2. 入出荷伝票参照の場合の価格決定要素データの取得  //2-2
			psdc.ItemPricingElement, e = f.DeliveryDocumentItemPricingElement(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 2-40. 入出荷伝票参照レコード・値の取得（入出荷伝票パートナ）  //I-2-1
			psdc.DeliveryDocumentPartner, e = f.DeliveryDocumentPartner(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				// 6-1-1. Orders Address からの住所データの取得  //I-2-1
				psdc.Address, e = f.DeliveryDocumentAddress(sdc, psdc)
				if e != nil {
					err = e
					return
				}

				// 6-2. AddressIDの登録(ユーザーが任意の住所を入力ファイルで指定した場合)
				psdc.Address, e = f.AddressFromInput(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)

			// InvoiceDocumentItem  //(I-1-2またはII-1-2)
			psdc.InvoiceDocumentItem = f.InvoiceDocumentItem(sdc, psdc)

			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				//3-1. InvoiceDocumentHeader  //I-1-1
				psdc.CalculateInvoiceDocument, e = f.CalculateInvoiceDocument(sdc, psdc)
				if e != nil {
					err = e
					return
				}
			}(wg)
		}(wg)

	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 99-1-2. CreationDate(Item)
		psdc.CreationDateItem = f.CreationDateItem(sdc, psdc)

		// 99-2-2. LastChangeDate(Item)
		psdc.LastChangeDateItem = f.LastChangeDateItem(sdc, psdc)

		// 99-3-2. CrationTime(Item)
		psdc.CreationTimeItem = f.CreationTimeItem(sdc, psdc)

		// 99-4-2. LastChangeTimeItem(Item)
		psdc.LastChangeTimeItem = f.LastChangeTimeItem(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}
