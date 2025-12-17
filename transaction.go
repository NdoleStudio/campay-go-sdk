package campay

// Transaction contains details of an initiated transaction
type Transaction struct {
	Reference         string `json:"reference"`
	Status            string `json:"status"`
	Amount            any    `json:"amount"`
	Currency          string `json:"currency"`
	Operator          string `json:"operator"`
	Code              string `json:"code"`
	OperatorReference string `json:"operator_reference"`
}

// IsPending checks if a transaction is pending
func (transaction *Transaction) IsPending() bool {
	return transaction.Status == "PENDING"
}

// IsSuccessful checks if a transaction is successful
func (transaction *Transaction) IsSuccessful() bool {
	return transaction.Status == "SUCCESSFUL"
}
