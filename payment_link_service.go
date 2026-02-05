package campay

import (
	"context"
	"encoding/json"
	"net/http"
)

// paymentLinkService is the API client for the `/get_payment_link/` endpoint
type paymentLinkService service

// Create payment links to receive payments from your clients using generated links.
//
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#f1b5aa4a-3146-42de-9250-f9b5406a7711
func (service *paymentLinkService) Create(ctx context.Context, params *PaymentLinkCreateRequest) (*PaymentLink, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/get_payment_link/", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	var paymentLink PaymentLink
	if err = json.Unmarshal(*response.Body, &paymentLink); err != nil {
		return nil, response, err
	}

	return &paymentLink, response, nil
}
