package campay

// Transaction contains details of an initiated transaction
type Transaction struct {
	Reference         string  `json:"reference"`
	Status            string  `json:"status"`
	Amount            float64 `json:"amount"`
	Currency          string  `json:"currency"`
	Operator          string  `json:"operator"`
	Code              string  `json:"code"`
	OperatorReference string  `json:"operator_reference"`
}

// IsPending checks if a transaction is pending
func (transaction *Transaction) IsPending() bool {
	return transaction.Status == "PENDING"
}

// IsSuccessfull checks if a transaction is successfull
func (transaction *Transaction) IsSuccessfull() bool {
	return transaction.Status == "SUCCESSFUL"
}
