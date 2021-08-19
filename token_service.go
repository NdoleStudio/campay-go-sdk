package campay

import (
	"context"
	"encoding/json"
	"net/http"
)

// TokenService is the API client for the `/gateway` endpoint
type TokenService service

// Get access token
// POST /token/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#8803168b-d451-4d65-b8cc-85e385bc3050
func (service *TokenService) Get(ctx context.Context) (*Token, *Response, error) {
	payload := map[string]string{
		"username": service.client.apiUsername,
		"password": service.client.apiPassword,
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/token", payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err := service.client.do(request)
	if err != nil {
		return nil, resp, err
	}

	var token Token
	if err = json.Unmarshal(*resp.Body, &token); err != nil {
		return nil, resp, err
	}

	return &token, resp, nil
}
