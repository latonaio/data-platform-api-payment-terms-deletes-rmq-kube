package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-payment-terms-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-payment-terms-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) PaymentTermsRead(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.PaymentTerms {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE paymentTerms.PaymentTerms = \"%s\" ", input.PaymentTerms.PaymentTerms),
		fmt.Sprintf("AND paymentTerms.BaseDate = %d ", input.PaymentTerms.BaseDate),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	paymentTerms.PaymentTerms,
    	paymentTerms.BaseDate,
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_payment_terms_payment_terms_data as paymentTerms 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToPaymentTerms(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
