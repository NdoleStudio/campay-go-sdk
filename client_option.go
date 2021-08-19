package campay

import (
	"net/http"
	"strings"
)

// ClientOption are options for constructing a client
type ClientOption interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithEnvironment sets the campay endpoint for API requests
// By default, ProdEnvironment is used.
func WithEnvironment(environment Environment) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		config.environment = Environment(strings.TrimRight(environment.String(), "/"))
	})
}

// WithAPIUsername sets the campay API username
func WithAPIUsername(apiUsername string) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiUsername = apiUsername
	})
}

// WithAPIPassword sets the campay API password
func WithAPIPassword(apiPassword string) ClientOption {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiPassword = apiPassword
	})
}
