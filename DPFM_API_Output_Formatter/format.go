package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToPriceMaster(rows *sql.Rows) (*PaymentTerms, error) {
	defer rows.Close()
	paymentTerms := PaymentTerms{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&paymentTerms.PaymentTerms,
			&paymentTerms.BaseDate,
			&paymentTerms.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &paymentTerms, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &paymentTerms, nil
	}

	return &paymentTerms, nil
}
