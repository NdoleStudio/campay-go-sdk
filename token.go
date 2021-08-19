package campay

// Token is the authentication token
type Token struct {
	Token     string `json:"token"`
	ExpiresIn int64  `json:"expires_in"`
}
