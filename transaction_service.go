package campay

import (
	"context"
	"encoding/json"
	"net/http"
)

// transactionService is the API client for the `/gateway` endpoint
type transactionService service

func (service *transactionService) Get(ctx context.Context, reference string) (*Transaction, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/transaction/"+reference, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var transaction Transaction
	if err = json.Unmarshal(*response.Body, &transaction); err != nil {
		return nil, response, err
	}

	return &transaction, response, nil
}
