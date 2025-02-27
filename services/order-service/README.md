## Folder Structure

```
order-service/
│── cmd/                  # Entry point (main.go)
│── infrastructure/       # Database, external dependencies
│── internal/
│   │── domain/          # Entities (Order Model)
│   │── handler/         # HTTP Handlers (OrderHandler)
│   │── repository/      # Database interactions
│   │── usecase/         # Business logic (OrderUseCase)
│── .dockerignore        # Ignore unnecessary files for Docker
│── .env                 # Environment variables
│── go.mod               # Module dependencies
│── go.sum               # Dependency lock file
```
