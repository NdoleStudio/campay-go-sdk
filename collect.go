package campay

// CollectParams is the details needed to initialize a payment
type CollectParams struct {
	Amount            uint   `json:"amount"`
	Currency          string `json:"currency"`
	From              string `json:"from"`
	Description       string `json:"description"`
	ExternalReference string `json:"external_reference"`
}

// CollectResponse is the response after calling the collect endpoint
type CollectResponse struct {
	Reference string `json:"reference"`
	UssdCode  string `json:"ussd_code"`
	Operator  string `json:"operator"`
}
