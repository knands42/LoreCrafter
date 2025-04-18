import typer
from rich import print

from src.character import generate_character
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


if __name__ == "__main__":
    app()
