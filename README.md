# Campay Go SDK

[![Build Status](https://travis-ci.com/NdoleStudio/campay-go-sdk.svg?branch=master)](https://travis-ci.com/NdoleStudio/campay-go-sdk)
[![codecov](https://codecov.io/gh/NdoleStudio/campay-go-sdk/branch/master/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/campay-go-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/campay-go-sdk)](https://goreportcard.com/report/github.com/NdoleStudio/campay-go-sdk)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/campay-go-sdk)](https://github.com/NdoleStudio/campay-go-sdk/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/campay-go-sdk?color=brightgreen)](https://github.com/NdoleStudio/campay-go-sdk/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/campay-go-sdk)](https://pkg.go.dev/github.com/NdoleStudio/campay-go-sdk)


This package provides a `go` client for interacting with the [CamPay REST API](https://documenter.getpostman.com/view/2391374/T1LV8PVA#intro)

## Installation

`campay-go-sdk` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/campay-go-sdk
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/campay-go-sdk"
```

## Implemented

- [Token](#token)
  - `POST /token` - Get access token
- [Collect](#collect)
  - `POST /collect` - Request Payment

## Usage

### Initializing the Client

An instance of the `campay` client can be created using `New()`.  The `http.Client` supplied will be used to make requests to the API.

```go
package main

import (
	"github.com/NdoleStudio/campay-go-sdk"
	"net/http"
)

func main()  {
	client := campay.New(
		campay.WithAPIUsername("" /* campay API Username */),
		campay.WithAPIPassword("" /* campay API Password */),
		campay.WithEnvironment(campay.DevEnvironment),
	)
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
payload, response, err := campayClient.Token(context.Background())
if err != nil {
  //handle error
}
```

### Token

This handles all API requests whose URL begins with `/token/`

#### Get access token

`POST /token/`: Get access token

```go
response, _, err := campayClient.Token(
	context.Background(),
)

if err != nil {
    log.Fatal(err)
}

log.Println(response.Token) // e.g eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsInVpZCI6Mn0.eyJpYXQiOjE2MDM4MjQ...
```

## Testing

You can run the unit tests for this SDK from the root directory using the command below:
```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
