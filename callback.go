package campay

// CallbackRequest is the notification request that is sent when a payment is made
type CallbackRequest struct {
	Status            string `json:"status"`
	Reference         string `json:"reference"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	Operator          string `json:"operator"`
	Code              string `json:"code"`
	OperatorReference string `json:"operator_reference"`
	Signature         string `json:"signature"`
	ExternalReference string `json:"external_reference"`
}
