# Clean Architecture Go Backend

A professional, production-ready Go backend template implementing Clean Architecture principles. Designed for high performance, maintainability, and clear separation of concerns.

---

## Project Overview
This project follows the **Clean Architecture** (Uncle Bob) pattern to ensure that the business logic remains independent of external frameworks, databases, and UI.

### Key Features
* **Clean Architecture Layers:** Strictly separated Entities, Services (Usecases), Repositories, and Handlers.
* **Database Migrations:** Version-controlled schema changes using `golang-migrate`.
* **Containerized:** Multi-stage `Dockerfile` and `docker-compose` for efficient deployment.
* **Modern Tooling:** Support for UUID primary keys, structured logging, and robust configuration management.
* **Linter Integration:** Pre-configured with `.golangci.yml` for code quality.

---

## Project Structure

```plaintext
├── cmd/
│   ├── api/            # Main entry point for the REST API
│   └── worker/         # Background workers/consumers
├── db/
│   └── migrations/     # SQL migration files (Up/Down)
├── internal/
│   ├── api/            # HTTP Handlers and Route definitions
│   ├── config/         # Database, Logger, and App configuration
│   ├── entity/         # Domain models (Business Logic Core)
│   ├── model/          # DTOs (Request/Response structs)
│   ├── repository/     # Data access layer (GORM/PostgreSQL)
│   └── services/       # Business logic implementation
├── config.json         # Local configuration file
├── docker-compose.yaml # Docker orchestration
└── Dockerfile          # Multi-stage build for Go
```

## Prerequisites
* Go (version 1.26+)
* Docker & Docker Compose
* PostgreSQL (if running locally without Docker)
* golang-migrate CLI (for creating new migrations)

## Getting Started
### 1. Clone the repository
```plaintext
git clone https://github.com/wayanedisudarma/golang-universe
cd clean-architecture
```
### 2. Configure Environment
Update your config.json or set environment variables for database connectivity. For Docker Compose, the default host is db-postgres.

### 3. Run with Docker Compose
The easiest way to start the app and database simultaneously:

```plaintext
docker-compose up
```

### 4. Running Locally (Native)
 ```plaintext
# Run the API
go run cmd/api/main.go
```

## Database Migrations
This project uses golang-migrate.

### Create a new migration:
```plaintext
migrate create -ext sql -dir db/migrations -seq <migration_name>
```
### Run migrations manually:
```plaintext
migrate -path db/migrations -database "postgres://user:pass@localhost:5432/clean_architecture?sslmode=disable" up
```

## Quality Control
### Linting:
Ensure your code adheres to the project standards before committing:
```plaintext
golangci-lint run
```

## Built With
* Gin Gonic - HTTP Web Framework
* GORM - ORM for Golang
* Viper - Configuration Management 
* Golang-Migrate - Database Migrations 
* Slog - Structured Logging