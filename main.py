import typer
from rich import print
from rich.prompt import Prompt
from character_generator import generate_character

app = typer.Typer()

def get_character_info():
    print("\n[bold cyan]Let's create your character![/bold cyan]")
    
    name = Prompt.ask("[bold green]What's your character's name?[/bold green]")
    
    race = Prompt.ask(
        "[bold green]Choose your character's race[/bold green]",
        choices=["Human", "Elf", "Dwarf", "Orc", "Halfling", "Other"],
        default="Human"
    )
    
    if race == "Other":
        race = Prompt.ask("[bold green]Please specify the race[/bold green]")
    
    personality = Prompt.ask(
        "[bold green]Describe your character's personality[/bold green]",
        default="Brave and adventurous"
    )
    
    return {
        "name": name,
        "race": race,
        "personality": personality
    }

@app.command()
def character():
    print("[bold green]🧙 Welcome to LoreCrafter![/bold green]")
    
    character_info = get_character_info()
    
    print("\n[bold yellow]Generating your character's story...[/bold yellow]")
    result = generate_character(character_info)
    print("\n[bold green]Your character's story:[/bold green]")
    print(result)

if __name__ == "__main__":
    app()
