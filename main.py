import json
from typing import Optional

import typer
from rich import print

from src.character import CharacterGenerator, CharacterVectorStore
from src.cli import get_character_info
from src.cli.helper import print_character

app = typer.Typer()

char_generator = CharacterGenerator()
char_vector_db = CharacterVectorStore()


@app.command()
def create_character():
    print("[bold cyan]🧙 Welcome to LoreCrafter![/bold cyan]")

    character_info = get_character_info()

    print("\n[bold yellow]Generating your character's story...[/bold yellow]")
    character_result = char_generator.generate(character_info)

    print_character(character_result)


@app.command()
def search_character(query: str, top: Optional[int] = 2):
    results = char_vector_db.search_similar(query, top)

    for i, doc in enumerate(results):
        print(f"\n--- Result #{i + 1} ---")
        retrieved_character = json.loads(doc.page_content)
        print_character(retrieved_character)


# @app.command()
# def create_character_with_reference(query: str):
#     results = retrieve_relevant_docs(query)
#     print(results)
# 

if __name__ == "__main__":
    app()
