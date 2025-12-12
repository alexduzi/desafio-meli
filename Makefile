.PHONY: help run build swagger test clean deps

# Default target
help:
	@echo "Available targets:"
	@echo "  make run        - Run the application locally"
	@echo "  make build      - Build the application binary"
	@echo "  make swagger    - Generate/regenerate Swagger documentation"
	@echo "  make test       - Run all tests with verbose output"
	@echo "  make clean      - Clean build artifacts and test cache"
	@echo "  make deps       - Download and tidy dependencies"
	@echo "  make all        - Run deps, swagger, build, and test"

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

# Run all tests with verbose output showing test names
test:
	@echo "Running all tests..."
	go test -v ./... -count=1

# Clean build artifacts and test cache
clean:
	@echo "Cleaning..."
	rm -rf bin/
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
