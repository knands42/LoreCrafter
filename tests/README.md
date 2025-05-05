# LoreCrafter Integration Tests

This directory contains integration tests for the LoreCrafter application, covering both the API and CLI interfaces.

## Test Structure

- `tests/api/`: Tests for the FastAPI endpoints
- `tests/cli/`: Tests for the Typer CLI commands
- `conftest.py`: Shared test fixtures and configuration

## Running Tests

You can run the tests using the following make commands:

```bash
# Run all tests
make test

# Run only API tests
make test-api

# Run only CLI tests
make test-cli

# Run tests with coverage reporting
make test-cov
```

## Test Coverage

The tests cover the following functionality:

### API Tests
- Basic endpoints (root, healthcheck)
- Character endpoints (create, search, get by ID)
- Error handling

### CLI Tests
- Character commands (create, search)
- World commands (create, search)
- Campaign commands (create)
- API server command

## Adding New Tests

When adding new tests:

1. Follow the existing patterns for mocking dependencies
2. Use fixtures from `conftest.py` where possible
3. Ensure tests are isolated and don't depend on external services
4. Add appropriate assertions to verify both function calls and return values

## Mocking Strategy

The tests use pytest's monkeypatch and unittest.mock to replace:
- External API calls (LLM services)
- Database operations (vector stores)
- File operations (PDF creation)
- User input (for CLI tests)

This ensures tests are fast, reliable, and don't require external dependencies.