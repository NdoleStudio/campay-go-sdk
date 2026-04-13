package campay

// PaymentOption represents a payment method available on the payment widget.
type PaymentOption string

var (
	// PaymentOptionMomo is the mobile money payment option.
	PaymentOptionMomo = PaymentOption("MOMO")

	// PaymentOptionCard is the card payment option.
	PaymentOptionCard = PaymentOption("CARD")
)

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
	Email              string `json:"email,omitempty"`
	FailureRedirectURL string `json:"failure_redirect_url,omitempty"`
	PaymentOptions     string `json:"payment_options,omitempty"`
}
