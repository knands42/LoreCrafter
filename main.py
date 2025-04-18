from typing import Optional

import typer
from rich import print
import json

from src.character import generate_character, search_similar_characters
from src.cli import get_character_info
from src.cli.helper import print_character

app = typer.Typer()


@app.command()
def character():
    print("[bold cyan]🧙 Welcome to LoreCrafter![/bold cyan]")

    character_info = get_character_info()
    
    print("\n[bold yellow]Generating your character's story...[/bold yellow]")
    character_result = generate_character(character_info)

    print_character(character_result)


@app.command()
def search_character(query: str, top: Optional[int] = 2):
    results = search_similar_characters(query, top)
    
    for i, doc in enumerate(results):
        print(f"\n--- Result #{i+1} ---")
        retrieved_character = json.loads(doc.page_content)
        print_character(retrieved_character)



if __name__ == "__main__":
    app()
