# Integration Tests

This directory contains integration tests for the LoreCrafter application. The tests are designed to verify the functionality of the API endpoints without mocking any dependencies.

## Test Structure

The tests follow the "Given, When, Then" pattern to make them more readable and maintainable:

- **Given**: The preconditions for the test
- **When**: The action being tested
- **Then**: The expected outcome

## Test Coverage

The integration tests cover the following features:

### Authentication

- User registration (success and failure scenarios)
- User login (success and failure scenarios)
- User profile retrieval (success and failure scenarios)

### Campaign Management

- Campaign creation (success and failure scenarios)
- Campaign retrieval (success and failure scenarios)
- Campaign update (success and failure scenarios)
- Campaign deletion (success and failure scenarios)
- Listing user campaigns (success and failure scenarios)

## Running the Tests

To run the integration tests, you need to have a PostgreSQL database running. The tests will use the following environment variables to connect to the database:

- `TEST_DB_HOST`: The database host (default: localhost)
- `TEST_DB_PORT`: The database port (default: 5432)
- `TEST_DB_USER`: The database user (default: postgres)
- `TEST_DB_PASSWORD`: The database password (default: postgres)
- `TEST_DB_NAME`: The database name (default: lorecrafter_test)

You can run the tests using the following command:

```bash
go test -v ./tests/integration
```

## Test Coverage Report

To generate a test coverage report, you can use the following command:

```bash
make test-coverage
```

This will generate a coverage report in HTML format (`coverage.html`) and display a summary of the coverage in the terminal. The coverage report excludes the `docs/*` and `pkg/*` directories as they are not part of the application's core functionality.