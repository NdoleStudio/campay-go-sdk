package campay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/campay-go-sdk/internal/helpers"
	"github.com/NdoleStudio/campay-go-sdk/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestPaymentLinkService_Create(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.PostTokenResponse(), stubs.CreatePaymentLinkResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(WithEnvironment(Environment(server.URL)))

	// Act
	link, response, err := client.PaymentLink.Create(context.Background(), &PaymentLinkCreateRequest{
		Amount:            "5",
		Currency:          "XAF",
		Description:       "Test",
		ExternalReference: "",
		RedirectURL:       "https://example.com/redirect",
	})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &PaymentLink{
		Link:      "http://127.0.0.1:8000/pay/test-xyz-1631658",
		Reference: "740871ff-c527-4474-be6b-147aaas2ea5",
	}, link)

	// Teardown
	server.Close()
}
