package campay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// holderInfoService is the API client for the `/holder_info/` endpoint
type holderInfoService service

// HolderInfo checks the name associated to a phone number
// GET /api/holder_info/?phone_number={phoneNumber}
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#7f8af623-82fd-4c17-850e-329d75823697
func (service *holderInfoService) HolderInfo(ctx context.Context, phoneNumber string) (*HolderInfo, *Response, error) {
	err := service.client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := service.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/holder_info/?phone_number=%s", phoneNumber), nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	holderInfo := new(HolderInfo)
	if err = json.Unmarshal(*response.Body, holderInfo); err != nil {
		return nil, response, err
	}

	return holderInfo, response, nil
}
