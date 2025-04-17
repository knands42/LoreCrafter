# LoreCrafter

LoreCrafter is an interactive CLI tool that helps you create rich backstories for your TTRPG characters. Using AI, it generates detailed narratives based on your character's basic information.

## Features

- Interactive character creation
- Customizable character attributes
- AI-generated detailed backstories
- Rich CLI interface with Typer

## Setup

1. Initialize the virtual environment (if not already done):
```bash
uv venv
```

2. Activate the virtual environment:
```bash
# On Windows
.venv\Scripts\activate
# On Unix/MacOS
source .venv/bin/activate
```

3. Install dependencies:
```bash
uv pip install -r pyproject.toml
```

4. Create a `.env` file with your OpenAI API key:
```bash
OPENAI_API_KEY=your_api_key_here
```

## Usage

Run the character generator:
```bash
python main.py character
```

Follow the interactive prompts to create your character's backstory.

## Development

- `make setup`: Initialize virtual environment and install dependencies
- `make run`: Run the character generator
- `make clean`: Remove virtual environment and cache files
