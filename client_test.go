package campay

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("default configuration is used when no option is set", func(t *testing.T) {
		// act
		client := New()

		// assert
		assert.NotEmpty(t, client.environment)
		assert.NotEmpty(t, client.common)

		assert.Empty(t, client.apiUsername)
		assert.Empty(t, client.apiPassword)

		assert.NotNil(t, client.httpClient)
		assert.NotNil(t, client.Transaction)
	})

	t.Run("single configuration value can be set using options", func(t *testing.T) {
		// Arrange
		env := Environment("https://example.com")

		// Act
		client := New(WithEnvironment(env))

		// Assert
		assert.NotNil(t, client.environment)
		assert.Equal(t, env.String(), client.environment.String())
	})

	t.Run("multiple configuration values can be set using options", func(t *testing.T) {
		// Arrange
		env := Environment("https://example.com")
		newHTTPClient := &http.Client{Timeout: 422}

		// Act
		client := New(WithEnvironment(env), WithHTTPClient(newHTTPClient))

		// Assert
		assert.NotEmpty(t, client.environment)
		assert.Equal(t, env.String(), client.environment.String())

		assert.NotNil(t, client.httpClient)
		assert.Equal(t, newHTTPClient.Timeout, client.httpClient.Timeout)
	})

	t.Run("it sets the Transaction service correctly", func(t *testing.T) {
		// Arrange
		client := New()

		// Assert
		assert.NotNil(t, client.Transaction)
		assert.Equal(t, client.environment.String(), client.Transaction.client.environment.String())
	})
}

func TestClient_ValidateCallback(t *testing.T) {
	t.Run("valid signature and key returns nil error", func(t *testing.T) {
		// Arrange
		client := New()
		key := []byte("geTrvNBLvXS35UvK3PnTnQgpjGmaEGe7wa6k3Ns4zehhvjncsjXnQsV7ZzhDWjDMEt7")
		signature := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.9kFRoXMiiGWhCbnDPoOMVhgVTzYxu2MjNi-uTWGRcEU"

		// Act
		err := client.ValidateCallback(signature, key)

		// Assert
		assert.Nil(t, err)
	})

	t.Run("invalid key returns and error", func(t *testing.T) {
		// Arrange
		client := New()
		key := []byte("geTrvNBLvXS35UvK3PnTnQgpjGmaEGe7wa6k3Ns4zehhvjncsjXnQsV7ZzhDWjDMEt7-invalid")
		signature := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.9kFRoXMiiGWhCbnDPoOMVhgVTzYxu2MjNi-uTWGRcEU"

		// Act
		err := client.ValidateCallback(signature, key)

		// Assert
		assert.NotNil(t, err)
	})
}
