# FIZZBUZZ-GO COMMAND RUNNER

# Default command
default:
    @just --list

# Run the server
start:
    go run .

# Execute unit tests
test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover ./...

# Build Docker image
build:
	@echo "Building Docker image..."
	docker build -t fizzbuzz-go:latest -f .docker/api.Dockerfile .
	@echo "Docker image built successfully!"

# Sync Go modules
tidy:
    cd api && go mod tidy
    go work sync

# Start up Docker Compose server
compose-up:
    @echo "Starting up full Docker suite..."
    docker compose build
    docker compose up

# Clean up Docker container remnants
compose-down:
    docker compose down
    @echo "Compose server closed successfully!"
