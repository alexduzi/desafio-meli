.PHONY: help run build swagger test test-unit test-integration test-coverage test-coverage-html clean deps docker-build docker-run docker-stop docker-logs docker-compose-up docker-compose-down docker-compose-logs docker-clean

# Default target
help:
	@echo "Available targets:"
	@echo ""
	@echo "Local Development:"
	@echo "  make run                 - Run the application locally"
	@echo "  make build               - Build the application binary"
	@echo "  make swagger             - Generate/regenerate Swagger documentation"
	@echo ""
	@echo "Testing:"
	@echo "  make test                - Run all tests (unit + integration)"
	@echo "  make test-unit           - Run only unit tests (fast, no DB required)"
	@echo "  make test-integration    - Run only integration tests (requires DB)"
	@echo "  make test-coverage       - Run tests with coverage report"
	@echo "  make test-coverage-html  - Generate HTML coverage report"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-build        - Build Docker image"
	@echo "  make docker-run          - Run Docker container"
	@echo "  make docker-stop         - Stop and remove Docker container"
	@echo "  make docker-logs         - View Docker container logs"
	@echo "  make docker-compose-up   - Start application with Docker Compose"
	@echo "  make docker-compose-down - Stop application with Docker Compose"
	@echo "  make docker-compose-logs - View Docker Compose logs"
	@echo "  make docker-clean        - Remove Docker images and containers"
	@echo ""
	@echo "Utilities:"
	@echo "  make clean               - Clean build artifacts and test cache"
	@echo "  make deps                - Download and tidy dependencies"
	@echo "  make all                 - Run deps, swagger, build, and test"

# Run the application locally
run:
	@echo "Starting application..."
	go run cmd/api/main.go

# Build the application
build:
	@echo "Building application..."
	go build -o bin/api cmd/api/main.go
	@echo "Build complete: bin/api"

# Generate/regenerate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@which swag > /dev/null || (echo "Installing swag..." && go install github.com/swaggo/swag/cmd/swag@latest)
	swag init -g cmd/api/main.go -o docs
	@echo "Swagger docs generated in docs/"

# Run all tests (unit + integration)
test:
	@echo "Running all tests..."
	go test -v ./... -count=1

# Run only unit tests (fast, no database required)
test-unit:
	@echo "Running unit tests..."
	go test -v -short ./internal/... -count=1

# Run only integration tests (requires database)
test-integration:
	@echo "Running integration tests..."
	go test -v ./test/integration/... -count=1

# Run tests with coverage report
test-coverage:
	@echo "Running tests with coverage..."
	@go test -short -cover ./internal/...
	@echo ""
	@echo "Coverage summary:"
	@go test -short -coverprofile=coverage.out ./internal/... > /dev/null 2>&1
	@go tool cover -func=coverage.out | grep total | awk '{print "Total coverage: " $$3}'
	@rm -f coverage.out

# Generate HTML coverage report and open in browser
test-coverage-html:
	@echo "Generating HTML coverage report..."
	@go test -short -coverprofile=coverage.out ./internal/...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@echo "Opening in browser..."
	@which xdg-open > /dev/null && xdg-open coverage.html || open coverage.html || echo "Please open coverage.html manually"

# Clean build artifacts and test cache
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	go clean -testcache
	@echo "Clean complete"

# Download and tidy dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "Dependencies updated"

# Run everything: deps, swagger, build, and test
all: deps swagger build test
	@echo "All tasks completed successfully!"

# Docker commands

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t product-api:latest .
	@echo "Docker image built successfully!"

# Run Docker container
docker-run:
	@echo "Starting Docker container..."
	docker run -d -p 8080:8080 --name product-api product-api:latest
	@echo "Container started! Access at http://localhost:8080"
	@echo "Health check: http://localhost:8080/health"
	@echo "Swagger UI: http://localhost:8080/swagger/index.html"

# Stop and remove Docker container
docker-stop:
	@echo "Stopping Docker container..."
	@docker stop product-api 2>/dev/null || true
	@docker rm product-api 2>/dev/null || true
	@echo "Container stopped and removed"

# View Docker container logs
docker-logs:
	@echo "Showing container logs (Ctrl+C to exit)..."
	docker logs -f product-api

# Start with Docker Compose
docker-compose-up:
	@echo "Starting application with Docker Compose..."
	docker-compose up -d
	@echo "Application started!"
	@echo "Access at http://localhost:8080"
	@echo "View logs: make docker-compose-logs"

# Stop Docker Compose
docker-compose-down:
	@echo "Stopping Docker Compose..."
	docker-compose down
	@echo "Application stopped"

# View Docker Compose logs
docker-compose-logs:
	@echo "Showing Docker Compose logs (Ctrl+C to exit)..."
	docker-compose logs -f

# Clean up Docker resources
docker-clean: docker-stop
	@echo "Cleaning up Docker resources..."
	@docker rmi product-api:latest 2>/dev/null || true
	@docker system prune -f
	@echo "Docker cleanup complete"
