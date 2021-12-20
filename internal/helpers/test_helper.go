package helpers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// MakeRequestCapturingTestServer creates an api server that captures the request object
func MakeRequestCapturingTestServer(responseCode int, responses [][]byte, requests *[]*http.Request) *httptest.Server {
	index := 0
	return httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		clonedRequest := request.Clone(context.Background())

		// clone body
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		request.Body = ioutil.NopCloser(bytes.NewReader(body))
		clonedRequest.Body = ioutil.NopCloser(bytes.NewReader(body))

		*requests = append(*requests, clonedRequest)

		responseWriter.WriteHeader(responseCode)
		_, err = responseWriter.Write(responses[index])
		if err != nil {
			panic(err)
		}
		index++
	}))
}
