package campay

// WithdrawParams are the details needed to perform a withdrawal
type WithdrawParams struct {
	Amount            uint    `json:"amount"`
	To                string  `json:"to"`
	Description       string  `json:"description"`
	ExternalReference *string `json:"external_reference,omitempty"`
}

// WithdrawResponse is the response after doing a withdrawal request
type WithdrawResponse struct {
	Reference string `json:"reference"`
	Status    string `json:"status"`
}
