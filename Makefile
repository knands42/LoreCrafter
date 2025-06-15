include .env.example

####### setup commands
# Generate Ed25519 key pair for PASETO tokens
generate-paseto-keys:
	openssl genpkey -algorithm Ed25519 -out private_key.pem
	openssl pkey -in private_key.pem -pubout -out public_key.pem

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
	swag init -g app/api/docs.go -o app/api/docs

.PHONY: sqlc-generate, swagger-generate, migrate-down, migrate-down1, migrate-up1, migration-up, migration-create, docker-build
