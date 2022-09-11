package campay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/campay-go-sdk/internal/helpers"
	"github.com/NdoleStudio/campay-go-sdk/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestUtilitiesService_AirtimeTransferSync(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.PostTokenResponse(), stubs.PostTransferResponse(), stubs.GetPendingAirtimeTransactionResponse(), stubs.GetSuccessfullAirtimeTransactionResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(WithEnvironment(Environment(server.URL)))

	// Act
	transaction, _, err := client.Utilities.AirtimeTransferSync(context.Background(), &AirtimeTransferParams{
		Amount:            "100",
		To:                "2376XXXXXXXXX",
		ExternalReference: "5577006791947779410",
	})

	// Assert
	assert.Nil(t, err)

	assert.GreaterOrEqual(t, len(requests), 4)
	assert.Equal(t, &UtilitiesTransaction{
		Reference:         "971e32ae-bb5a-420a-a38a-c2931536609f",
		ExternalReference: "5577006791947779410",
		Status:            "SUCCESSFUL",
		Amount:            100,
		Currency:          "XAF",
		Operator:          "ORANGE_CM",
		Code:              "CP220804U0649K",
		Type:              "AIRTIME",
		Reason:            "",
	}, transaction)

	assert.True(t, transaction.IsSuccessfull())
	assert.False(t, transaction.IsPending())

	// Teardown
	server.Close()
}
