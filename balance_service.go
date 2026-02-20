package campay

import (
	"context"
	"encoding/json"
	"net/http"
)

// balanceService is the API client for the `/balance/` endpoint
type balanceService service

// Get Application Balance
//
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#7e2b34aa-6645-4d75-abe1-c28499f9d34d
func (service *balanceService) Get(ctx context.Context) (*Balance, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, "/api/balance/", nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var balance Balance
	if err = json.Unmarshal(*response.Body, &balance); err != nil {
		return nil, response, err
	}

	return &balance, response, nil
}
