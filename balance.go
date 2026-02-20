package campay

// Balance is the response from the `/balance/` endpoint
type Balance struct {
	TotalBalance  int    `json:"total_balance"`
	MtnBalance    int    `json:"mtn_balance"`
	OrangeBalance int    `json:"orange_balance"`
	Currency      string `json:"currency"`
}
