from dotenv import load_dotenv

from src.adapter.input.cli import CharCLIShell, WorldCLIShell, CampaignCLIShell, print_character, print_world, \
    print_campaign
from src.adapter.output.llm import LLMFactory
from src.adapter.output.pdf import create_character_pdf, create_world_pdf
from src.adapter.output.repository import CharacterVectorStore, WorldVectorStore
from src.application.usecases import CharacterGenerator, WorldGenerator, CampaignGenerator

load_dotenv()

import json
from typing import Optional

import typer
from rich import print
from rich.console import Console

app = typer.Typer()
llm = LLMFactory.create()
console = Console()

char_vector_db = CharacterVectorStore()
char_generator = CharacterGenerator(char_vector_db)
world_vector_db = WorldVectorStore()
world_generator = WorldGenerator(world_vector_db)
campaign_generator = CampaignGenerator(world_vector_db)

charCLIShell = CharCLIShell(console, world_vector_db)
worldCLIShell = WorldCLIShell(console, world_vector_db)
campaignCLIShell = CampaignCLIShell(console, world_vector_db)


@app.command()
def create_character(
    default: bool = typer.Option(False, "--default", help="Use predefined character information")
):
    print("[bold cyan]🧙 Welcome to LoreCrafter![/bold cyan]")

    character_info = charCLIShell.get_character_info(default)

    print("\n[bold yellow]Generating your character's story...[/bold yellow]")
    character_result = char_generator.generate(character_info)

    print_character(character_result)
    create_character_pdf(character_result, "character_profile.pdf")


@app.command()
def create_world(
    default: bool = typer.Option(False, "--default", help="Use predefined world information")
):
    print("[bold cyan]🌍 Welcome to LoreCrafter World Builder![/bold cyan]")

    world_info = worldCLIShell.get_world_info(default)

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


@app.command()
def create_campaign(
    default: bool = typer.Option(False, "--default", help="Use predefined campaign information")
):
    print("[bold cyan]🎭 Welcome to LoreCrafter Campaign Builder![/bold cyan]")

    campaign_info = campaignCLIShell.get_campaign_info(default)

    print("\n[bold yellow]Generating your campaign...[/bold yellow]")
    campaign_result = campaign_generator.generate(campaign_info)

    print_campaign(campaign_result)


if __name__ == "__main__":
    app()
