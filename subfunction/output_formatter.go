package subfunction

import (
	api_input_reader "data-platform-api-invoice-document-items-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-invoice-document-items-creates-subfunc/API_Processing_Data_Formatter"
)

func (f *SubFunction) SetValue(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	item, err := dpfm_api_output_formatter.ConvertToItem(sdc, psdc)
	if err != nil {
		return err
	}

	itemPricingElement, err := dpfm_api_output_formatter.ConvertToItemPricingElement(sdc, psdc)
	if err != nil {
		return err
	}

	partner, err := dpfm_api_output_formatter.ConvertToPartner(sdc, psdc)
	if err != nil {
		return err
	}

	address, err := dpfm_api_output_formatter.ConvertToAddress(sdc, psdc)
	if err != nil {
		return err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Item:               item,
		ItemPricingElement: itemPricingElement,
		Partner:            partner,
		Address:            address,
	}

	return nil
}
