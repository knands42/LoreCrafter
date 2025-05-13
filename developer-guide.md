# Developers Guide

## Features

- Generate character backstories, personalities, and appearances
- Create world histories and timelines
- Design campaign settings and hidden elements
- Generate character and world images
- Export character and world information as PDFs
- Search for characters and worlds

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/LoreCrafter.git
cd LoreCrafter
```

2. Install the required dependencies and set up the environment:
```bash
make create-env
make install-deps
make generate-keys
```

3. Set up environment variables:
Create a `.env` file in the root directory with your API keys and other configuration.

## Usage

### CLI Mode

LoreCrafter can be used as a command-line tool:

```bash
# Create a character
python main.py create-character

# Create a world
python main.py create-world

# Create a campaign
python main.py create-campaign

# Search for a character
python main.py search-character "brave warrior"

# Search for a world
python main.py search-world "fantasy kingdom"
```

### API Mode

LoreCrafter also provides a RESTful API using FastAPI:

```bash
# Start the API server
python main.py api

# Start with custom host and port
python main.py api --host 0.0.0.0 --port 8080

# Start with auto-reload for development
python main.py api --reload
```

Once the API server is running, you can access the API documentation at `http://localhost:8000/docs`.

## API Documentation

The API provides endpoints for creating and retrieving characters, worlds, and campaigns. See the [API Documentation](frontend/docs/api_documentation.md) for details.

## Project Structure

- `src/adapter/input/cli`: CLI interface
- `src/adapter/input/api`: API interface
- `src/adapter/output/repository`: Data storage
- `src/adapter/output/llm`: LLM integration
- `src/adapter/output/pdf`: PDF generation
- `src/application/domain`: Domain models
- `src/application/usecases`: Application logic
- `src/application/prompts`: LLM prompts
- `tests/api`: API integration tests
- `tests/cli`: CLI integration tests

## Testing

LoreCrafter includes comprehensive integration tests for both the API and CLI interfaces. You can run the tests using the following commands:

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

For more details about the testing approach, see the [Tests README](tests/README.md).

## API Documentation

LoreCrafter provides a RESTful API with comprehensive documentation using Swagger UI and ReDoc.

### Accessing the API Documentation

Once the application is running, you can access the API documentation at:

- **Swagger UI**: [http://localhost:8000/docs](http://localhost:8000/docs)
- **ReDoc**: [http://localhost:8000/redoc](http://localhost:8000/redoc)

### API Endpoints

The API is organized into the following categories:

- **Characters**: Create, retrieve, and search for characters with AI-generated backstories
- **Worlds**: Create, retrieve, and search for worlds with AI-generated histories and timelines
- **Campaigns**: Create and retrieve campaigns with AI-generated settings and hidden elements
- **Assets**: Retrieve character and world images
- **PDF**: Generate PDF documents for characters and worlds

## License

This project is licensed under the MIT License - see the LICENSE file for details.
