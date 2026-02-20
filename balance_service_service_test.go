package campay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/campay-go-sdk/internal/helpers"
	"github.com/NdoleStudio/campay-go-sdk/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestBalanceService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.PostTokenResponse(), stubs.BalanceResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(WithEnvironment(Environment(server.URL)))

	// Act
	link, response, err := client.Balance.Get(context.Background())

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &Balance{
		TotalBalance:  3,
		MtnBalance:    2,
		OrangeBalance: 1,
		Currency:      "XAF",
	}, link)

	// Teardown
	server.Close()
}
