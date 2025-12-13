.PHONY: help run build swagger test test-unit test-integration test-coverage test-coverage-html clean deps

# Default target
help:
	@echo "Available targets:"
	@echo "  make run                 - Run the application locally"
	@echo "  make build               - Build the application binary"
	@echo "  make swagger             - Generate/regenerate Swagger documentation"
	@echo "  make test                - Run all tests (unit + integration)"
	@echo "  make test-unit           - Run only unit tests (fast, no DB required)"
	@echo "  make test-integration    - Run only integration tests (requires DB)"
	@echo "  make test-coverage       - Run tests with coverage report"
	@echo "  make test-coverage-html  - Generate HTML coverage report"
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
