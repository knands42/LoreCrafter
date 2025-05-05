from dotenv import load_dotenv

load_dotenv()

from src.adapter.input.cli import app

if __name__ == "__main__":
    app()
