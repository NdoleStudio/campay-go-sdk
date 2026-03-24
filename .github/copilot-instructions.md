# Copilot Instructions for campay-go-sdk

## Build, Test, and Lint

```bash
# Run all tests with race detection
go test -v -race ./...

# Run a single test
go test -v -run TestClient_Withdraw ./...

# Lint (requires golangci-lint)
golangci-lint run

# Format and tidy
go-fumpt -w .
go mod tidy
```

## Architecture

This is a Go SDK for the [CamPay](https://www.campay.net) mobile money API. The client uses a **service-oriented composition pattern**:

- `Client` is the entry point, created via `New(...ClientOption)` using the functional options pattern.
- Services (`balanceService`, `transactionService`, `paymentLinkService`, `utilitiesService`) are unexported types that alias a shared `service` struct embedding a `*Client`.
- Services are wired during `New()` via pointer casting: `client.Transaction = (*transactionService)(&client.common)`.
- Token refresh is thread-safe (`sync.Mutex`) and lazily triggered before each API call via `service.client.refreshToken(ctx)`.

### Request lifecycle

Every API method follows the same flow:

1. `service.client.refreshToken(ctx)` — acquire/reuse JWT token
2. `service.client.newRequest(ctx, method, path, body)` — build `*http.Request` with JSON body and `Authorization: Token` header
3. `service.client.do(req)` — execute and return `(*Response, error)`
4. `json.Unmarshal(*response.Body, &target)` — decode into domain struct

### Return convention

All public API methods return `(T, *Response, error)`. `Response` wraps `*http.Response` and the raw body, and has an `Error()` method that returns non-nil for non-2xx status codes.

### Sync vs Async operations

Some operations have both variants (e.g., `Withdraw` / `WithdrawSync`, `AirtimeTransfer` / `AirtimeTransferSync`). The `*Sync` variants poll the transaction status every 10 seconds for up to ~5 minutes until the transaction leaves `PENDING` state.

## Key Conventions

### Adding a new API endpoint

1. Create the request/response model types in a new `<name>.go` file with `json` struct tags.
2. Create `<name>_service.go` with an unexported service type aliasing `service` (e.g., `type fooService service`).
3. Add the service as a public field on `Client` and wire it in `New()`.
4. Follow the existing request lifecycle pattern (refreshToken → newRequest → do → Unmarshal).

### Testing

- Tests use `github.com/stretchr/testify/assert` — no test suites or table-driven tests.
- All tests run in parallel (`t.Parallel()`).
- HTTP calls are mocked with `internal/helpers.MakeRequestCapturingTestServer`, which returns a `*httptest.Server` that captures requests and returns canned responses in sequence.
- Canned JSON responses live in `internal/stubs/api_responses.go` as `func() []byte`.
- Tests follow AAA with explicit `// Arrange`, `// Act`, `// Assert`, `// Teardown` comments.
- Test naming: `Test<Type>_<Method>` (e.g., `TestBalanceService_Get`).
- Since the first API call is always token acquisition, stub arrays start with `stubs.PostTokenResponse()`.

### JSON encoding

- `json.NewEncoder` with `SetEscapeHTML(false)` for request bodies.
- Optional fields use pointer types with `omitempty` (e.g., `ExternalReference *string`).

### Environments

`Environment` is a string type with two predefined values: `DevEnvironment` (`https://demo.campay.net`) and `ProdEnvironment` (`https://www.campay.net`). Tests override this by passing the test server URL as an `Environment`.
