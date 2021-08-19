package campay

import (
	"net/http"
)

type clientConfig struct {
	httpClient  *http.Client
	environment Environment
	apiPassword string
	apiUsername string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient:  http.DefaultClient,
		apiPassword: "",
		apiUsername: "",
		environment: ProdEnvironment,
	}
}
