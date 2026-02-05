package campay

// PaymentLink contains data for Campay payment links
type PaymentLink struct {
	Link      string `json:"link"`
	Reference string `json:"reference"`
}

// PaymentLinkCreateRequest creates payment links to receive payments from your clients.
type PaymentLinkCreateRequest struct {
	Amount             string `json:"amount"`
	Currency           string `json:"currency"`
	Description        string `json:"description"`
	ExternalReference  string `json:"external_reference"`
	RedirectURL        string `json:"redirect_url"`
	FailureRedirectURL string `json:"failure_redirect_url,omitempty"`
	PaymentOptions     string `json:"payment_options,omitempty"`
}
