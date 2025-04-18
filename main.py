import typer
from rich import print
from src.character import generate_character
from src.cli import get_character_info

app = typer.Typer()


@app.command()
def character():
    print("[bold green]🧙 Welcome to LoreCrafter![/bold green]")
    
    character_info = get_character_info()
    custom_story = character_info.pop("custom_story", None)
    
    print("\n[bold yellow]Generating your character's story...[/bold yellow]")
    result = generate_character(character_info, custom_story)
    print("\n[bold green]Your character's story:[/bold green]")
    print(result)

if __name__ == "__main__":
    app()
