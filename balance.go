package campay

// Balance is the response from the `/balance/` endpoint
type Balance struct {
	TotalBalance  float64 `json:"total_balance"`
	MtnBalance    float64 `json:"mtn_balance"`
	OrangeBalance float64 `json:"orange_balance"`
	Currency      string  `json:"currency"`
}
