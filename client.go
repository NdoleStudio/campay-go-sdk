package campay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/golang-jwt/jwt"
)

type service struct {
	client *Client
}

// Client is the campay API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient          *http.Client
	common              service
	environment         Environment
	mutex               sync.Mutex
	apiUsername         string
	apiPassword         string
	token               string
	tokenExpirationTime int64
	Transaction         *transactionService
}

// New creates and returns a new campay.Client from a slice of campay.ClientOption.
func New(options ...ClientOption) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		httpClient:  config.httpClient,
		environment: config.environment,
		apiUsername: config.apiUsername,
		apiPassword: config.apiPassword,
		mutex:       sync.Mutex{},
	}

	client.common.client = client
	client.Transaction = (*transactionService)(&client.common)
	return client
}

// Token Gets the access token
// POST /token/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#8803168b-d451-4d65-b8cc-85e385bc3050
func (client *Client) Token(ctx context.Context) (*Token, *Response, error) {
	payload := map[string]string{
		"username": client.apiUsername,
		"password": client.apiPassword,
	}

	request, err := client.newRequest(ctx, http.MethodPost, "/token/", payload)
	if err != nil {
		return nil, nil, err
	}

	resp, err := client.do(request)
	if err != nil {
		return nil, resp, err
	}

	var token Token
	if err = json.Unmarshal(*resp.Body, &token); err != nil {
		return nil, resp, err
	}

	return &token, resp, nil
}

// ValidateCallback checks if the signature was encrypted with the webhook key
func (client *Client) ValidateCallback(signature string, webhookKey []byte) error {
	_, err := jwt.Parse(signature, func(token *jwt.Token) (interface{}, error) {
		return webhookKey, nil
	})
	return err
}

// Collect Requests a Payment
// POST /collect/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#31757962-2e07-486b-a6f4-a7cc7a06d032
func (client *Client) Collect(ctx context.Context, params *CollectParams) (*CollectResponse, *Response, error) {
	err := client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := client.newRequest(ctx, http.MethodPost, "/collect/", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := client.do(request)
	if err != nil {
		return nil, response, err
	}

	var collectResponse CollectResponse
	if err = json.Unmarshal(*response.Body, &collectResponse); err != nil {
		return nil, response, err
	}

	return &collectResponse, response, nil
}

// Withdraw funds to a mobile money account
// POST /withdraw/
// API Doc: https://documenter.getpostman.com/view/2391374/T1LV8PVA#885dbde0-b0dd-4514-a0f9-f84fc83df12d
func (client *Client) Withdraw(ctx context.Context, params *WithdrawParams) (*WithdrawResponse, *Response, error) {
	err := client.refreshToken(ctx)
	if err != nil {
		return nil, nil, err
	}

	request, err := client.newRequest(ctx, http.MethodPost, "/withdraw/", params)
	if err != nil {
		return nil, nil, err
	}

	response, err := client.do(request)
	if err != nil {
		return nil, response, err
	}

	var withdrawResponse WithdrawResponse
	if err = json.Unmarshal(*response.Body, &withdrawResponse); err != nil {
		return nil, response, err
	}

	return &withdrawResponse, response, nil
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, client.environment.String()+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if len(client.token) > 0 {
		req.Header.Set("Authorization", "Token "+client.token)
	}

	return req, nil
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("%T cannot be nil", req)
	}

	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(ioutil.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// refreshToken refreshes the authentication Token
func (client *Client) refreshToken(ctx context.Context) error {
	if client.tokenExpirationTime > time.Now().UTC().Unix() {
		return nil
	}

	client.mutex.Lock()
	defer client.mutex.Unlock()

	client.token = ""

	token, _, err := client.Token(ctx)
	if err != nil {
		return err
	}

	client.token = token.Token
	client.tokenExpirationTime = time.Now().UTC().Unix() + token.ExpiresIn - 1000 // Give extra 100 second buffer

	return nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	if httpResponse == nil {
		return nil, fmt.Errorf("%T cannot be nil", httpResponse)
	}

	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := ioutil.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	return resp, resp.Error()
}
