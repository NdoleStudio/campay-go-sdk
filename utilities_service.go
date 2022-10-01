package campay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
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

// AirtimeTransferSync transfers airtime to a mobile number and waits for the transaction to be completed.
// POST /api/utilities/airtime/transfer/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#544cb091-1104-4bf1-b9ad-893c5067c925
func (service *utilitiesService) AirtimeTransferSync(ctx context.Context, params *AirtimeTransferParams) (*UtilitiesTransaction, *Response, error) {
	transaction, response, err := service.AirtimeTransfer(ctx, params)
	if err != nil {
		return nil, response, err
	}

	// wait for completion in 2 minutes
	counter := 1
	for {
		status, response, err := service.TransactionStatus(ctx, transaction.Reference)
		if err != nil || !status.IsPending() || ctx.Err() != nil || counter == 30 {
			return status, response, err
		}
		time.Sleep(10 * time.Second)
		counter++
	}
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
