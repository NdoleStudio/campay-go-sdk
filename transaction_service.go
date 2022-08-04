package campay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// transactionService is the API client for the `/transaction/` endpoint
type transactionService service

// Get the status of an initiated transaction.
// GET /transaction/{reference}/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#d9278e44-aa8a-4ed1-85f7-7e0184dc35db
func (service *transactionService) Get(ctx context.Context, reference string) (*Transaction, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/transaction/%s/", reference), nil)
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
