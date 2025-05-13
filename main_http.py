from dotenv import load_dotenv

load_dotenv()

from src.adapter.input.api.app import app

if __name__ == "__main__":
    app()
