# Order Service - Microservices

## Overview

Order Service is a microservice responsible for handling orders in the MicroShop system. It follows **Clean Architecture** principles and uses **Golang, Fiber, GORM, and PostgreSQL**.

## Tech Stack

- **Language:** Golang
- **Framework:** Fiber
- **Database:** PostgreSQL + GORM
- **Dependency Management:** Go Modules
- **Configuration:** `.env` using `godotenv`
- **Containerization:** Docker

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
