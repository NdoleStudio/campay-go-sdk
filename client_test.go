package campay

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/NdoleStudio/campay-go-sdk/internal/helpers"
	"github.com/NdoleStudio/campay-go-sdk/internal/stubs"

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

func TestClient_Withdraw(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.PostTokenResponse(), stubs.PostWithdrawResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(WithEnvironment(Environment(server.URL)))

	// Act
	withdrawResponse, response, err := client.Withdraw(context.Background(), &WithdrawParams{
		Amount:            100,
		To:                "2376XXXXXXXX",
		Description:       "Test",
		ExternalReference: nil,
	})

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, "/api/withdraw/", request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Token"))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &WithdrawResponse{Reference: "26676007-1c31-46d7-9c71-acb031cf0de4", Status: "PENDING"}, withdrawResponse)

	// Teardown
	server.Close()
}

func TestClient_WithdrawSync(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.PostTokenResponse(), stubs.PostWithdrawResponse(), stubs.GetPendingTransactionResponse(), stubs.GetSuccessfulTransactionResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(WithEnvironment(Environment(server.URL)))

	// Act
	transaction, _, err := client.WithdrawSync(context.Background(), &WithdrawParams{
		Amount:            100,
		To:                "2376XXXXXXXX",
		Description:       "Test",
		ExternalReference: nil,
	})

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 4)
	assert.Equal(t, &Transaction{
		Reference:         "bcedde9b-62a7-4421-96ac-2e6179552a1a",
		Status:            "SUCCESSFUL",
		Amount:            1,
		Currency:          "XAF",
		Operator:          "MTN",
		Code:              "CP201027T00005",
		OperatorReference: "1880106956",
	}, transaction)

	assert.True(t, transaction.IsSuccessful())
	assert.False(t, transaction.IsPending())

	// Teardown
	server.Close()
}
