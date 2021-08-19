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
	Token               *TokenService
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
	client.Token = (*TokenService)(&client.common)
	return client
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (c *Client) newRequest(ctx context.Context, method, uri string, body interface{}) (*http.Request, error) {
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

	req, err := http.NewRequestWithContext(ctx, method, c.environment.String()+uri, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if len(c.token) > 0 {
		req.Header.Set("Authorization", "Token "+c.token)
	}

	return req, nil
}

// do carries out an HTTP request and returns a Response
func (c *Client) do(req *http.Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("%T cannot be nil", req)
	}

	httpResponse, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := c.newResponse(httpResponse)
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
func (c *Client) refreshToken(ctx context.Context) error {
	if c.tokenExpirationTime > time.Now().UTC().Unix() {
		return nil
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	token, _, err := c.Token.Get(ctx)
	if err != nil {
		return err
	}

	c.tokenExpirationTime = time.Now().UTC().Unix() + token.ExpiresIn - 100 // Give extra 100 second buffer

	return nil
}

// newResponse converts an *http.Response to *Response
func (c *Client) newResponse(httpResponse *http.Response) (*Response, error) {
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
