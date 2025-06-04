# Generate Ed25519 key pair for PASETO tokens
generate-keys:
	openssl genpkey -algorithm Ed25519 -out private_key.pem
	openssl pkey -in private_key.pem -pubout -out public_key.pem

# Build the application
build:
	go build -o bin/lorecrafter .

# Run the application
run: build
	./bin/lorecrafter

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Docker commands
docker-build:
	docker build -t lorecrafter:latest .

docker-run:
	docker run -p 8080:8080 --env-file .env lorecrafter:latest

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

docker-compose-logs:
	docker-compose logs -f

# Database migration commands
migrate-create:
	@read -p "Enter migration name: " name; \
	mkdir -p migrations; \
	touch migrations/$(shell date +%Y%m%d%H%M%S)_$$name.up.sql; \
	touch migrations/$(shell date +%Y%m%d%H%M%S)_$$name.down.sql

migrate-up:
	@echo "Running migrations up..."
	@echo "This is a placeholder. Replace with your actual migration command."
	# Example: migrate -path=./migrations -database=postgres://postgres:postgres@localhost:5432/lorecrafter?sslmode=disable up

migrate-down:
	@echo "Running migrations down..."
	@echo "This is a placeholder. Replace with your actual migration command."
	# Example: migrate -path=./migrations -database=postgres://postgres:postgres@localhost:5432/lorecrafter?sslmode=disable down 1

# Generate SQLC code
sqlc-generate:
	sqlc generate

# Setup development environment
setup: generate-keys
	@echo "Setting up development environment..."
	go mod download
	@echo "Development environment setup complete."


.PHONY: generate-keys build run test clean docker-build docker-run docker-compose-up docker-compose-down docker-compose-logs migrate-create migrate-up migrate-down sqlc-generate setup help
