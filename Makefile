include .env.example

SHELL := /bin/bash

# SSL/TLS configuration
SSL_DIR := certs
KEY_FILE := $(SSL_DIR)/server.key
CERT_FILE := $(SSL_DIR)/server.crt

####### setup commands

# Generate self-signed certificate for local development
setup-ssl:
	@echo "Generating self-signed certificate..."
	@mkdir -p $(SSL_DIR)
	@openssl req -x509 \
		-nodes \
		-days 365 \
		-newkey rsa:2048 \
		-keyout $(KEY_FILE) \
		-out $(CERT_FILE) \
		-subj "/C=US/ST=State/L=City/O=Company/CN=lorecrafter.fly.dev" \
		-addext "subjectAltName=DNS:DNS:lorecrafter.fly.dev,IP:127.0.0.1"
	@echo "Self-signed certificate generated in $(SSL_DIR)/"
	@echo "Base64 encoding keys and updating .env file..."
	@if [ ! -f .env ]; then cp .env.example .env; fi
	@SSL_CERT_BASE64=$$(cat $(CERT_FILE) | base64 -w 0) && \
	SSL_KEY_BASE64=$$(cat $(KEY_FILE) | base64 -w 0) && \
	sed -i "s|SSL_CERT=.*|SSL_CERT=$$SSL_CERT_BASE64|" .env && \
	sed -i "s|SSL_KEY=.*|SSL_KEY=$$SSL_KEY_BASE64|" .env
	@echo "Keys generated, encoded, and set in .env file"

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
setup: setup-paseto-keys setup-password-salt setup-ssl

####### application commands #######
# Build the application
build:
	go build -o bin/lorecrafter .

# Run the application with HTTP/2 and TLS
run: build
	@echo "Starting server with HTTP/2 and TLS on https://localhost:$(SERVER_PORT)"
	@echo "Note: You may need to accept the self-signed certificate in your browser"
	./bin/lorecrafter --tls

# Run the application without TLS (HTTP only)
run-http: build
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

####### docker commands #######
docker-build:
	docker rmi lorecrafter || true
	docker build -t lorecrafter:latest .

docker-run:
	docker run -p 8000:8000 --env-file .env lorecrafter:latest

docker-up:
	docker rmi lorecrafter || true
	$(MAKE) docker-build
	docker-compose up --build --force-recreate

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

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

.PHONY: sqlc-generate, swagger-generate, migration-down, migration-down1, migration-up1, migration-up, migration-create, docker-build, setup-ssl
