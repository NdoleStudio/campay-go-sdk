package campay

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("httpClient is not set when the httpClient is nil", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()

		// Act
		WithHTTPClient(nil).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
	})

	t.Run("httpClient is set when the httpClient is not nil", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()
		newClient := &http.Client{Timeout: 300}

		// Act
		WithHTTPClient(newClient).apply(config)

		// Assert
		assert.NotNil(t, config.httpClient)
		assert.Equal(t, newClient.Timeout, config.httpClient.Timeout)
	})
}

func TestWithEnvironment(t *testing.T) {
	t.Run("environment is set successfully", func(t *testing.T) {
		// Arrange
		environment := Environment("https://example.com")
		config := defaultClientConfig()

		// Act
		WithEnvironment(environment).apply(config)

		// Assert
		assert.NotNil(t, config.environment)
		assert.Equal(t, environment.String(), config.environment.String())
	})

	t.Run("tailing / is trimmed from environment", func(t *testing.T) {
		// Arrange
		environment := Environment("https://example.com/")
		config := defaultClientConfig()

		// Act
		WithEnvironment(environment).apply(config)

		// Assert
		assert.NotNil(t, config.environment)
		assert.Equal(t, "https://example.com", config.environment.String())
	})
}

func TestWithAPIUsername(t *testing.T) {
	t.Run("apiUser is set successfully", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()
		apiUser := "apiUser"

		// Act
		WithAPIUsername(apiUser).apply(config)

		// Assert
		assert.Equal(t, apiUser, config.apiUsername)
	})
}

func TestWithAPIPassword(t *testing.T) {
	t.Run("apiPassword is set successfully", func(t *testing.T) {
		// Arrange
		config := defaultClientConfig()
		apiPassword := "apiPassword"

		// Act
		WithAPIPassword(apiPassword).apply(config)

		// Assert
		assert.Equal(t, apiPassword, config.apiPassword)
	})
}
