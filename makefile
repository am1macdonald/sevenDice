.PHONY: dev build test clean migrate-up migrate-down sqlc

# Development
dev:
	docker compose up --build

# Build the application
build:
	go build -o bin/server ./cmd/server

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/ tmp/

# Database migrations
migrate-up:
	goose -dir sql/migrations postgres "postgres://sevendice:sevendice@localhost:5432/sevendice?sslmode=disable" up

migrate-down:
	goose -dir sql/migrations postgres "postgres://sevendice:sevendice@localhost:5432/sevendice?sslmode=disable" down

# Generate SQLC
sqlc:
	sqlc generate

# Install development tools
tools:
	go install github.com/air-verse/air@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

# Development initialization
init: tools
	go mod tidy
	make migrate-up
	make sqlc
