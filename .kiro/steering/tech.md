# LoreCrafter Technology Stack

## Core Technologies

- **Language**: Go 1.23+
- **Web Framework**: Chi router v5 for HTTP routing and middleware
- **Database**: PostgreSQL with pgx/v5 driver
- **Query Builder**: SQLC for type-safe SQL code generation
- **Authentication**: PASETO tokens with Ed25519 keys
- **Password Hashing**: Argon2 for secure password storage
- **Configuration**: Viper for environment-based config management
- **Containerization**: Docker/Podman with multi-stage builds

## AI & External Services

- **LLM Integration**: LangChain Go for multiple AI providers (OpenAI, Google, Anthropic)
- **Email Service**: Mailgun for transactional emails
- **Caching/Queue**: Valkey (Redis-compatible) for background jobs
- **Documentation**: Swagger/OpenAPI with auto-generation

## Development Tools

- **Database Migrations**: golang-migrate for schema versioning
- **Testing**: Standard Go testing with integration test suite
- **Code Coverage**: Built-in Go coverage tools with HTML reports
- **API Documentation**: Swaggo for automatic Swagger generation
- **Git Hooks**: Pre-commit hooks for code quality

## Common Commands

### Development
```bash
make setup              # Initial project setup (keys, env, hooks)
make build              # Build the application binary
make run                # Build and run the server
make test               # Run all tests
make test-coverage      # Run tests with coverage report
```

### Database Operations
```bash
make migration-create NAME=<name>  # Create new migration
make migration-up                  # Apply all pending migrations
make migration-down1               # Rollback one migration
make sqlc-generate                 # Generate type-safe SQL code
```

### Container Operations
```bash
make podman-build       # Build container image
make podman-up          # Start all services
make podman-up-dbs      # Start only databases
make podman-down        # Stop all services
```

### Documentation
```bash
make swagger-generate   # Generate API documentation
```

## Architecture Patterns

- **Clean Architecture**: Domain-driven design with clear separation of concerns
- **Dependency Injection**: Constructor-based DI for testability
- **Repository Pattern**: Database abstraction through interfaces
- **Use Case Pattern**: Business logic encapsulation in dedicated use cases