# Integration Tests

This directory contains integration tests for the LoreCrafter application. The tests are designed to verify the functionality of the API endpoints without mocking any dependencies.

## Test Structure

The tests follow the "Given, When, Then" pattern to make them more readable and maintainable:

- **Given**: The preconditions for the test
- **When**: The action being tested
- **Then**: The expected outcome

### Test Cleanup

Each test automatically cleans up after itself using Go's `t.Cleanup()` mechanism. This ensures that:

- Data created during a test is removed after the test completes
- Tests are isolated from each other
- The database remains clean between test runs

The cleanup process:
1. Tracks users created during tests
2. Automatically deletes these users after each test
3. Cascades deletions to related data (campaigns, campaign members, etc.) through database constraints

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

```bash
make test
```

## Test Coverage Report

To generate a test coverage report, you can use the following command:

```bash
make test-coverage
```

This will generate a coverage report in HTML format (`coverage.html`) and display a summary of the coverage in the terminal. The coverage report excludes the `docs/*` and `pkg/*` directories as they are not part of the application's core functionality.
