from dotenv import load_dotenv

from src.content_generator.pdf_generator import create_character_pdf, create_world_pdf

load_dotenv()

import json
from typing import Optional

import typer
from rich import print

from src.character import CharacterGenerator, CharacterVectorStore
from src.world import WorldGenerator
from src.world.world_vector_store import WorldVectorStore
from src.cli import get_character_info
from src.cli.get_world_info import get_world_info
from src.cli.helper import print_character, print_world

app = typer.Typer()

char_generator = CharacterGenerator()
char_vector_db = CharacterVectorStore()
world_generator = WorldGenerator()
world_vector_db = WorldVectorStore()


@app.command()
def create_character(
    default: bool = typer.Option(False, "--default", help="Use predefined character information")
):
    print("[bold cyan]🧙 Welcome to LoreCrafter![/bold cyan]")

    character_info = get_character_info(default)

    print("\n[bold yellow]Generating your character's story...[/bold yellow]")
    character_result = char_generator.generate(character_info)

    print_character(character_result)
    create_character_pdf(character_result, "character_profile.pdf")


@app.command()
def create_world(
    default: bool = typer.Option(False, "--default", help="Use predefined world information")
):
    print("[bold cyan]🌍 Welcome to LoreCrafter World Builder![/bold cyan]")

    world_info = get_world_info(default)

    print("\n[bold yellow]Generating your world...[/bold yellow]")
    world_result = world_generator.generate(world_info)

    print_world(world_result)
    create_world_pdf(world_result, "world_lore.pdf")


@app.command()
def search_character(query: str, top: Optional[int] = 2):
    results = char_vector_db.search_similar(query, top)

    for i, doc in enumerate(results):
        print(f"\n--- Result #{i + 1} ---")
        retrieved_character = json.loads(doc.page_content)
        print_character(retrieved_character)


@app.command()
def search_world(query: str, top: Optional[int] = 2):
    results = world_vector_db.search_similar(query, top)

    for i, doc in enumerate(results):
        print(f"\n--- Result #{i + 1} ---")
        retrieved_world = json.loads(doc.page_content)
        print_world(retrieved_world)


if __name__ == "__main__":
    app()
