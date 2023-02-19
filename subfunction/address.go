package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) OrdersAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.Address, error) {
	args := make([]interface{}, 0)

	ordersHeader := psdc.OrdersHeader
	repeat := strings.Repeat("?,", len(ordersHeader)-1) + "?"
	for _, v := range ordersHeader {
		args = append(args, v.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, AddressID, PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_address_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersAddress(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) DeliveryDocumentAddress(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.Address, error) {
	args := make([]interface{}, 0)

	deliveryDocumentHeader := psdc.DeliveryDocumentHeader
	repeat := strings.Repeat("?,", len(deliveryDocumentHeader)-1) + "?"
	for _, v := range deliveryDocumentHeader {
		args = append(args, v.DeliveryDocument)
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, AddressID, PostalCode, LocalRegion, Country, District, StreetName, CityName, Building, Floor, Room
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_address_data
		WHERE DeliveryDocument IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToDeliveryDocumentAddress(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) AddressFromInput(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.Address, error) {
	processFlag := false

	calculateAddressID, err := f.CalculateAddressID(sdc, psdc)
	if err != nil {
		return nil, err
	}

	addressMasterdata := make([]*api_processing_data_formatter.AddressMaster, 0)
	addressID := calculateAddressID.AddressID
	for i, v := range sdc.Header.Address {
		if v.PostalCode != nil || v.LocalRegion != nil || v.Country != nil || v.District != nil || v.StreetName != nil || v.CityName != nil {
			if len(*v.PostalCode) != 0 || len(*v.LocalRegion) != 0 || len(*v.Country) != 0 || len(*v.District) != 0 || len(*v.StreetName) != 0 || len(*v.CityName) != 0 {
				processFlag = true
				datum := psdc.ConvertToAddressMaster(sdc, i, addressID)
				addressMasterdata = append(addressMasterdata, datum)
				addressID += 1
			}
		}
	}
	psdc.AddressMaster = addressMasterdata

	if !processFlag {
		return psdc.Address, nil
	}

	sessionID := sdc.RuntimeSessionID
	for _, addressData := range addressMasterdata {
		res, err := f.rmq.SessionKeepRequest(f.ctx, f.conf.RMQ.QueueToSQL(), map[string]interface{}{"message": addressData, "function": "Address", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			f.l.Error(err)
			return []*api_processing_data_formatter.Address{}, nil
		}
		res.Success()
	}

	data := psdc.ConvertToAddressFromInput()

	return data, err
}

func (f *SubFunction) CalculateAddressID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateAddressID, error) {
	dataKey := psdc.ConvertToCalculateAddressIDKey()

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dataQueryGets, err := psdc.ConvertToCalculateAddressIDQueryGets(rows)
	if err != nil {
		return nil, err
	}

	if dataQueryGets.LatestNumber == nil {
		return nil, xerrors.Errorf("LatestNumberがnullです。")
	}

	addressIDLatestNumber := dataQueryGets.LatestNumber
	addressID := *dataQueryGets.LatestNumber + 1

	data := psdc.ConvertToCalculateAddressID(addressIDLatestNumber, addressID)

	return data, err
}
