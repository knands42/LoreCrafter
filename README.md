# LoreCrafter

LoreCrafter is a tool that helps users create rich backstories for tabletop role-playing game (TTRPG) characters, worlds, and campaigns using AI. The system generates detailed narratives based on user-provided information and can link characters, worlds, and campaigns together to create a cohesive storytelling experience.

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

2. Install the required dependencies:
```bash
uv pip install .
uv pip install .[test]
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

## License

This project is licensed under the MIT License - see the LICENSE file for details.
