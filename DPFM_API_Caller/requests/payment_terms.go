package requests

type PaymentTerms struct {
	PaymentTerms        string `json:"PaymentTerms"`
	BaseDate            int    `json:"BaseDate"`
	IsMarkedForDeletion *bool  `json:"IsMarkedForDeletion"`
}
