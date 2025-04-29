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

# Run fuzz tests
fuzz path:
    @echo "Running fuzz tests!"
    go test -fuzz=Fuzz {{ path }} -fuzztime=30s

# Run fuzz tests
benchmark path:
    @echo "Running fuzz tests!"
    go test -bench=. -benchmem  {{ path }}

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
