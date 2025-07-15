include .env.example

SHELL := /bin/bash

####### setup commands

# Generate Ed25519 key pair for PASETO tokens
setup-paseto-keys:
	openssl genpkey -algorithm Ed25519 -out private_key.pem
	openssl pkey -in private_key.pem -pubout -out public_key.pem
	@echo "Base64 encoding keys and updating .env file..."
	@if [ ! -f .env ]; then cp .env.example .env; fi
	@PRIVATE_KEY_BASE64=$$(cat private_key.pem | base64 -w 0) && \
	PUBLIC_KEY_BASE64=$$(cat public_key.pem | base64 -w 0) && \
	sed -i "s|PASETO_PRIVATE_KEY=.*|PASETO_PRIVATE_KEY=$$PRIVATE_KEY_BASE64|" .env && \
	sed -i "s|PASETO_PUBLIC_KEY=.*|PASETO_PUBLIC_KEY=$$PUBLIC_KEY_BASE64|" .env
	@echo "Keys generated, encoded, and set in .env file"
	@echo "Cleaning up temporary PEM files..."
	@rm private_key.pem public_key.pem

setup-password-salt:
	@echo "Generating password salt..."
	@PASSWORD_SALT=$$(head -c 16 /dev/urandom | base64) && \
	sed -i "s|PASSWORD_SALT=.*|PASSWORD_SALT=$$PASSWORD_SALT|" .env
	@echo "Password salt generated and set in .env file"

# Setup all required configurations
setup: setup-paseto-keys setup-password-salt

####### application commands #######
# Build the application
build:
	go build -o bin/lorecrafter .

# Run the application
run: build
	@echo "Starting server without TLS on http://localhost:$(SERVER_PORT)"
	./bin/lorecrafter

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out -covermode=atomic $(shell go list ./... | grep -v "docs\|pkg")
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out

# Clean build artifacts
clean:
	rm -rf bin/

####### podman commands #######
podman-build:
	podman rmi lorecrafter || true
	podman build -t lorecrafter:latest .

podman-run:
	podman run -p 8000:8000 --env-file .env lorecrafter:latest

podman-up:
	podman rmi lorecrafter || true
	$(MAKE) podman-build
	podman compose up --build --force-recreate

podman-down:
	podman compose down

podman-logs:
	podman compose logs -f

####### migration commands #######
# e.g., make migration-create NAME=create-users
migration-create:
	migrate create -ext sql -dir internal/adapter/database/migrations -seq $(NAME)

migration-up:
	migrate -path internal/adapter/database/migrations -database "$(POSTGRES_URL)" up

migration-up1:
	migrate -path internal/adapter/database/migrations -database "$(POSTGRES_URL)" -verbose up 1

migration-down:
	migrate -path internal/adapter/database/migrations -database "$(POSTGRES_URL)" -verbose down

migration-down1:
	migrate -path internal/adapter/database/migrations -database "$(POSTGRES_URL)" -verbose down 1

# Generate SQLC code
sqlc-generate:
	sqlc generate

# Generate Swagger documentation
swagger-generate:
	swag init -g app/api/docs.go -o app/api/docs --parseDependency

.PHONY: sqlc-generate, swagger-generate, migration-down, migration-down1, migration-up1, migration-up, migration-create, podman-build, setup-ssl, test
