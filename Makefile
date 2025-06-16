include .env.example

SHELL := /bin/bash

####### setup commands
# Generate Ed25519 key pair for PASETO tokens
setup:
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

####### application commands #######
# Build the application
build:
	go build -o bin/lorecrafter .

# Run the application
run: build
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

migrate-up1:
	migrate -path internal/adapter/database/migrations -database "$(POSTGRES_URL)" -verbose up 1

migrate-down:
	migrate -path internal/adapter/database/migrations -database "$(POSTGRES_URL)" -verbose down

migrate-down1:
	migrate -path internal/adapter/database/migrations -database "$(POSTGRES_URL)" -verbose down 1

# Generate SQLC code
sqlc-generate:
	sqlc generate

# Generate Swagger documentation
swagger-generate:
	swag init -g app/api/docs.go -o app/api/docs --parseDependency

.PHONY: sqlc-generate, swagger-generate, migrate-down, migrate-down1, migrate-up1, migration-up, migration-create, docker-build
