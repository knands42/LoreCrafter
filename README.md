# LoreCrafter - UserCreation Authentication Service

A Go-based user authentication service implementing clean architecture principles. This service provides user registration and authentication using Argon2 for password hashing and PASETO tokens for authentication.

## Architecture

The project follows clean architecture principles with the following layers:

- **Domain**: Contains the core business entities and interfaces
- **Use Cases**: Implements the application business rules
- **Interfaces**: Contains the HTTP handlers and other adapters
- **Infrastructure**: Contains technical details like database connections and security utilities
- **Repositories**: Implements data access logic

## Features

- UserCreation registration with email and username
- UserCreation authentication with username and password
- Password hashing using Argon2id
- Token-based authentication using PASETO (Platform-Agnostic Security Tokens)
- Clean architecture implementation
- PostgreSQL database with SQLC for type-safe SQL

## Project Structure

```
.
├── internal/
│   ├── domain/           # Core business entities and interfaces
│   ├── usecases/         # Application business rules
│   ├── interfaces/       # HTTP handlers and adapters
│   ├── repositories/     # Data access implementations
│   └── infrastructure/   # Technical details (DB, security, etc.)
├── main.go               # Application entry point
└── sqlc.yaml             # SQLC configuration
```

## Setup Instructions

### Prerequisites

- Go 1.23 or higher
- PostgreSQL database
- SQLC (for generating database code)

### Configuration

The main configuration is in `main.go`:

```go
// Configuration constants
const (
    ServerPort     = "8080"
    TokenExpiry    = 24 * time.Hour
    PrivateKeyPath = "private_key.pem"
    PublicKeyPath  = "public_key.pem"
)
```

For a production environment, these should be moved to environment variables or a configuration file.

### Key Generation

Before running the application, you need to generate the Ed25519 key pair for PASETO token signing:

```bash
make generate-keys
```

This will create `private_key.pem` and `public_key.pem` files in the project root.

### Database Setup

1. Create a PostgreSQL database
2. Run the SQL migrations:

```bash
make migrate-up
```

3. Generate the SQLC code:

```bash
make sqlc-generate
```

### Running the Application

You can run the application using the Makefile:

```bash
make run
```

Or with Docker Compose:

```bash
make docker-compose-up
```

The server will start on port 8080 by default.

## API Documentation

### Authentication Endpoints

#### Register a New UserCreation

```
POST /api/auth/register
```

Request body:
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

Response:
```json
{
  "user": {
    "id": "uuid-string",
    "username": "johndoe",
    "email": "john@example.com"
  },
  "token": "paseto-token-string",
  "expires_at": "expiry-date-time"
}
```

#### Login

```
POST /api/auth/login
```

Request body:
```json
{
  "username": "johndoe",
  "password": "securepassword"
}
```

Response:
```json
{
  "user": {
    "id": "uuid-string",
    "username": "johndoe",
    "email": "john@example.com"
  },
  "token": "paseto-token-string",
  "expires_at": "expiry-date-time"
}
```

### Protected Endpoints

To access protected endpoints, include the token in the Authorization header:

```
Authorization: Bearer paseto-token-string
```

Example protected endpoint:

```
GET /api/me
```

## Security Considerations

- The PASETO token private key should be kept secret and never shared
- The public key can be shared with frontend clients for token verification
- In production, use environment variables or a secure configuration management system for key paths
- Consider adding rate limiting to prevent brute force attacks
- Use HTTPS in production

## Makefile Commands

The project includes a Makefile with various commands to simplify development:

```
make generate-keys     - Generate Ed25519 key pair for PASETO tokens
make build             - Build the application
make run               - Run the application
make test              - Run tests
make clean             - Clean build artifacts
make docker-build      - Build Docker image
make docker-run        - Run Docker container
make docker-compose-up - Start services with Docker Compose
make docker-compose-down - Stop services with Docker Compose
make docker-compose-logs - View Docker Compose logs
make migrate-create    - Create a new migration
make migrate-up        - Run migrations up
make migrate-down      - Run migrations down
make sqlc-generate     - Generate SQLC code
make setup             - Setup development environment
make help              - Show this help message
```

Run `make help` to see this list of commands.
