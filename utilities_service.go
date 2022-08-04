package campay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// utilitiesService is the API client for the `/transaction/` endpoint
type utilitiesService service

// AirtimeTransfer transfers airtime to a mobile number
// POST /api/utilities/airtime/transfer/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#544cb091-1104-4bf1-b9ad-893c5067c925
func (service *utilitiesService) AirtimeTransfer(ctx context.Context, params *AirtimeTransferParams) (*AirtimeTransferResponse, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodPost, "/api/utilities/airtime/transfer/", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(AirtimeTransferResponse)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}

// TransactionStatus checks the status of an initiated utility (Airtime) transaction
// GET /api/utilities/transaction/(reference)/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#e2686c97-565d-45dd-86f4-375e39268a44
func (service *utilitiesService) TransactionStatus(ctx context.Context, reference string) (*UtilitiesTransaction, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/utilities/transaction/%s/", reference), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	transaction := new(UtilitiesTransaction)
	if err = json.Unmarshal(*response.Body, transaction); err != nil {
		return nil, response, err
	}

	return transaction, response, nil
}
