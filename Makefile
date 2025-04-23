.PHONY: build run test migrate seed clean

# Default variables
APP_NAME := hospital-portal
PORT := 8000

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) cmd/server/main.go

# Run the application

run:
	@echo "Running $(APP_NAME)..."
	@go run cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# Run database migrations up
migrate-up:
	@echo "Running migrations up..."
	@bash scripts/migrate.sh up

# Run database migrations down
migrate-down:
	@echo "Running migrations down..."
	@bash scripts/migrate.sh down

# Create a new migration
migrate-create:
	@echo "Creating new migration..."
	@bash scripts/migrate.sh create $(name)

# Seed the database
seed:
	@echo "Seeding database..."
	@bash scripts/seed.sh

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@go clean

# Format the code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run linters
lint:
	@echo "Running linters..."
	@golangci-lint run

# Help
help:
	@echo "Make commands for $(APP_NAME):"
	@echo "  build           - Build the application"
	@echo "  run             - Run the application"
	@echo "  test            - Run tests"
	@echo "  test-coverage   - Run tests with coverage"
	@echo "  migrate-up      - Run database migrations up"
	@echo "  migrate-down    - Run database migrations down"
	@echo "  migrate-create  - Create a new migration (usage: make migrate-create name=migration_name)"
	@echo "  seed            - Seed the database with sample data"
	@echo "  clean           - Clean build artifacts"
	@echo "  fmt             - Format the code"
	@echo "  lint            - Run linters"
