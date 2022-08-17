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

// IsPending checks if a transaction is pending
func (transaction *UtilitiesTransaction) IsPending() bool {
	return transaction.Status == "PENDING"
}

// IsSuccessfull checks if a transaction is successfull
func (transaction *UtilitiesTransaction) IsSuccessfull() bool {
	return transaction.Status == "SUCCESSFUL"
}
