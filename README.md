# Campay Go SDK

[![Build](https://github.com/NdoleStudio/campay-go-sdk/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/campay-go-sdk/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/campay-go-sdk/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/campay-go-sdk)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/campay-go-sdk/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/campay-go-sdk/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/campay-go-sdk)](https://goreportcard.com/report/github.com/NdoleStudio/campay-go-sdk)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/campay-go-sdk)](https://github.com/NdoleStudio/campay-go-sdk/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/campay-go-sdk?color=brightgreen)](https://github.com/NdoleStudio/campay-go-sdk/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/campay-go-sdk)](https://pkg.go.dev/github.com/NdoleStudio/campay-go-sdk)


This package provides a `go` client for interacting with the [CamPay API](https://documenter.getpostman.com/view/2391374/T1LV8PVA#intro)

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
  - `POST /token`: Get access token
- [Collect](#collect)
  - `POST /collect`: Request Payment
- [Withdraw](#withdraw)
  - `POST /withdraw`: Withdraw funds to a mobile money account
- [Transaction](#transaction)
  - `GET /transaction/{reference}/`: Get the status of a transaction
- [Utilities](#utilities)
  - `POST /api/utilities/airtime/transfer/`: Transfers airtime to a mobile number
  - `GET /api/utilities/transaction/{reference}/`: Get the status of a transaction

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
token, _, err := campayClient.Token(context.Background())

if err != nil {
    log.Fatal(err)
}

log.Println(token.Token) // e.g eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsInVpZCI6Mn0...
```

### Collect

This handles all API requests whose URL begins with `/collect/`

#### `POST /collect/`: Request Payment

This endpoint is used to request payment from users.

```go
collectResponse, httpResponse, err := campayClient.Collect(context.Background(), campay.CollectOptions{
    Amount: 100,
    Currency: "XAF",
    From: "2376XXXXXXXX",
    Description: "Test",
    ExternalReference: "",
})

if err != nil {
    log.Fatal(err)
}

log.Prinln(collectResponse.Reference) // e.g 26676007-1c31-46d7-9c71-acb031cf0de4
```

### Withdraw

This handles all API requests whose URL begins with `/withdraw/`

#### `POST /withdraw`: Withdraw funds to a mobile money account

Withdraw funds from an app to a mobile money account.

```go
withdrawResponse, response, err := client.Withdraw(context.Background(), &WithdrawParams{
    Amount:            100,
    To:                "2376XXXXXXXX",
    Description:       "Test",
    ExternalReference: nil,
})

if err != nil {
    log.Fatal(err)
}

log.Println(withdrawResponse.Reference) // e.g 26676007-1c31-46d7-9c71-acb031cf0de4
```

### Transaction

This handles all API requests whose URL begins with `/transaction/`

#### `GET /transaction/{reference}/`: Get the status of a transaction

Use this endpoint to check for the status of an initiated transaction.

```go
transaction, httpResponse, err := campayClient.Transaction.Get(
	context.Background(),
	"bcedde9b-62a7-4421-96ac-2e6179552a1a"
)

if err != nil {
    log.Fatal(err)
}

log.Println(transaction.Reference) // e.g 26676007-1c31-46d7-9c71-acb031cf0de4
```

### Utilities

#### `POST /api/utilities/airtime/transfer/`: Transfers airtime to a mobile number

```go
transaction, httpResponse, err := client.Utilities.AirtimeTransferSync(context.Background(), &AirtimeTransferParams{
  Amount:            "100",
  To:                "237677777777",
  ExternalReference: "sample-external-ref",
})

if err != nil {
    log.Fatal(err)
}

log.Println(transaction.Reference) // e.g 26676007-1c31-46d7-9c71-acb031cf0de4
```

#### `GET /api/utilities/transaction/{reference}/`: Get the status of a transaction

```go
transaction, httpResponse, err := client.Utilities.AirtimeTransferSync(context.Background(), "" /* Transaction reference */)
if err != nil {
    log.Fatal(err)
}

log.Println(transaction.Status) // e.g "SUCCESSFUL"
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

### Security

If you discover any security related issues, please email arnoldewin@gmail.com instead of using the GitHub issues.

## Credits

- [Acho Arnold](https://github.com/AchoArnold)
- [All Contributors](../../contributors)


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
