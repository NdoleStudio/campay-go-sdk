package campay

// CallbackRequest is the notification request that is sent when a payment is made
type CallbackRequest struct {
	Status            string `json:"status" query:"status"`
	Reference         string `json:"reference" query:"reference"`
	Amount            string `json:"amount" query:"amount"`
	Currency          string `json:"currency" query:"currency"`
	Operator          string `json:"operator" query:"operator"`
	Code              string `json:"code" query:"code"`
	OperatorReference string `json:"operator_reference" query:"operator_reference"`
	Signature         string `json:"signature" query:"signature"`
	ExternalReference string `json:"external_reference" query:"external_reference"`
}
