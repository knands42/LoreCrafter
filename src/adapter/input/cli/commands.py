import json
from typing import Optional

import typer
import uvicorn
from rich import print
from rich.console import Console

from src.adapter.input.cli.helper import print_character, print_world, print_campaign
from src.adapter.input.cli.user_input.campaign_cli_shell import CampaignCLIShell
from src.adapter.input.cli.user_input.char_cli_shell import CharCLIShell
from src.adapter.input.cli.user_input.world_cli_shell import WorldCLIShell
from src.adapter.output.llm import LLMFactory
from src.adapter.output.repository import CharacterVectorStore, WorldVectorStore
from src.application.usecases import CharacterGenerator, WorldGenerator, CampaignGenerator

app = typer.Typer()
llm = LLMFactory.create_chat()
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


@app.command()
def create_world(
    default: bool = typer.Option(False, "--default", help="Use predefined world information")
):
    print("[bold cyan]🌍 Welcome to LoreCrafter World Builder![/bold cyan]")

    world_info = worldCLIShell.get_world_info(default)

    print("\n[bold yellow]Generating your world...[/bold yellow]")
    world_result = world_generator.generate(world_info)

    print_world(world_result)


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


@app.command()
def api(
    host: str = "127.0.0.1",
    port: int = 8000,
    reload: bool = False
):
    """
    Run the FastAPI server for the LoreCrafter API.
    """
    print("[bold cyan]🚀 Starting LoreCrafter API server...[/bold cyan]")

    # Import here to avoid circular imports

    uvicorn.run(
        "src.adapter.input.api.app:app",
        host=host,
        port=port,
        reload=reload
    )
