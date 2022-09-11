package campay

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/NdoleStudio/campay-go-sdk/internal/helpers"
	"github.com/NdoleStudio/campay-go-sdk/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestTransactionService_Get(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.PostTokenResponse(), stubs.GetPendingTransactionResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(WithEnvironment(Environment(server.URL)))
	reference := "bcedde9b-62a7-4421-96ac-2e6179552a1a"

	// Act
	withdrawResponse, response, err := client.Transaction.Get(context.Background(), reference)

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 1)
	request := requests[len(requests)-1]
	assert.Equal(t, fmt.Sprintf("/api/transaction/%s/", reference), request.URL.Path)
	assert.True(t, strings.HasPrefix(request.Header.Get("Authorization"), "Token"))
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &Transaction{
		Reference:         "bcedde9b-62a7-4421-96ac-2e6179552a1a",
		Status:            "PENDING",
		Amount:            1,
		Currency:          "XAF",
		Operator:          "MTN",
		Code:              "CP201027T00005",
		OperatorReference: "1880106956",
	}, withdrawResponse)

	// Teardown
	server.Close()
}
