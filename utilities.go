package campay

// AirtimeTransferParams are parameters for transferring airtime
type AirtimeTransferParams struct {
	Amount            string `json:"amount"`
	To                string `json:"to"`
	ExternalReference string `json:"external_reference"`
}

// AirtimeTransferResponse is gotten after transferring airtime
type AirtimeTransferResponse struct {
	Reference string `json:"reference"`
	Status    string `json:"status"`
}

// UtilitiesTransaction represent a utility transaction
type UtilitiesTransaction struct {
	Reference         string      `json:"reference"`
	ExternalReference string      `json:"external_reference"`
	Status            string      `json:"status"`
	Amount            float64     `json:"amount"`
	Currency          string      `json:"currency"`
	Operator          string      `json:"operator"`
	Code              string      `json:"code"`
	Type              string      `json:"type"`
	Reason            interface{} `json:"reason"`
}
