package campay

import (
	"context"
	"net/http"
	"testing"

	"github.com/NdoleStudio/campay-go-sdk/internal/helpers"
	"github.com/NdoleStudio/campay-go-sdk/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestHolderInfoService_HolderInfo(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	requests := make([]*http.Request, 0)
	responses := [][]byte{stubs.PostTokenResponse(), stubs.GetHolderInfoResponse()}
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, responses, &requests)
	client := New(WithEnvironment(Environment(server.URL)))

	// Act
	holderInfo, response, err := client.HolderInfo.HolderInfo(context.Background(), "237XXXXXXXX")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, response.HTTPResponse.StatusCode)

	assert.Equal(t, &HolderInfo{
		FullName: "JOHN DOE",
	}, holderInfo)

	// Teardown
	server.Close()
}
