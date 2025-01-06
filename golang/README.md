# README.md

# Gift Service

This project is a simple implementation of a gift evaluation service, where Santa evaluates gift requests based on children's behavior and the feasibility of the gifts.

## Project Structure

```
gift-service
├── cmd
│   └── main.go                # Entry point of the application
├── internal
│   └── gift
│       ├── behavior.go        # Defines the Behavior type
│       ├── child.go           # Defines the Child struct
│       ├── gift_request.go     # Defines the GiftRequest struct
│       ├── priority.go        # Defines the Priority type
│       └── santa_service.go   # Defines the SantaService struct
├── test
│   └── gift
│       └── santa_service_test.go # Unit tests for SantaService
├── go.mod                     # Module definition
├── go.sum                     # Module dependency checksums
└── README.md                  # Project documentation
```

## Getting Started

To run the application, use the following command:

```
go run cmd/main.go
```

## Testing

To run the tests, use the following command:

```
go test ./...
```

## License

This project is licensed under the MIT License.