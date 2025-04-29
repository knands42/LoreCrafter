import re
from typing import List, Optional

from rich import print
from rich.prompt import Prompt, Confirm
from rich.console import Console
from rich.table import Table

from src.world.world_vector_store import WorldVectorStore


def clean_text(value: str) -> str:
    # remove all characters except a-z, 0-9, and spaces
    return re.sub(r'[^a-zA-Z0-9 ]+', '', value)


def multiline_input(prompt: str) -> str:
    """Capture multiline input until the user presses Enter twice."""
    print(f"\n[bold cyan]{prompt}[/bold cyan]")
    lines = []
    while True:
        line = input()
        if not line and lines and not lines[-1]:
            break
        lines.append(line)
    return "\n".join(lines[:-1])  # Remove the last empty line


def display_and_select_world() -> Optional[dict]:
    """Display a list of available worlds and let the user select one.

    Returns:
        The selected world as a dictionary, or None if no world is selected.
    """
    world_store = WorldVectorStore()
    worlds = world_store.get_all_worlds()

    if not worlds:
        print("[yellow]No worlds found. You'll need to create a character without linking to a world.[/yellow]")
        return None

    console = Console()
    table = Table(title="Available Worlds")

    table.add_column("#", style="dim")
    table.add_column("Name", style="bold")
    table.add_column("Universe", style="cyan")
    table.add_column("Theme", style="green")

    for i, world in enumerate(worlds, 1):
        table.add_row(
            str(i),
            world.get("name", "Unknown"),
            world.get("universe", "Unknown"),
            world.get("world_theme", "Unknown")
        )

    console.print(table)

    # Ask user to select a world
    selection = Prompt.ask(
        "[bold green]Select a world by number (or press Enter to skip)[/bold green]",
        default=""
    )

    if not selection:
        return None

    try:
        index = int(selection) - 1
        if 0 <= index < len(worlds):
            return worlds[index]
        else:
            print("[red]Invalid selection. Creating character without linking to a world.[/red]")
            return None
    except ValueError:
        print("[red]Invalid input. Creating character without linking to a world.[/red]")
        return None


def ask_with_examples(prompt_text: str, examples: list[str], default: str = None) -> str:
    example_str = f"[dim]e.g., {', '.join(examples)}[/dim]"
    full_prompt = f"{prompt_text} {example_str}"
    return Prompt.ask(full_prompt, default=default)


def get_character_info(get_default: bool = False):
    if get_default:
        return {
            "name": "Captain Kirk",
            "gender": "male",
            "race": "Human",
            "personality": None,
            "appearance": None,
            "universe": "D&D",
            "world_theme": "fantasy",
            "tone": "Epic",
            "custom_story": None,
            "linked_world_id": None
        }

    print("\n[bold magenta]Let's create your character![/bold magenta]")

    # Ask if the user wants to link with an existing world
    link_with_world = Confirm.ask(
        "[bold yellow]Do you want to link your character with an existing world?[/bold yellow]",
        default=False
    )

    linked_world = None
    universe = None
    theme = None

    if link_with_world:
        linked_world = display_and_select_world()

        if linked_world:
            # Use universe and world_theme from the linked world
            universe = linked_world.get("universe", "D&D")
            theme = linked_world.get("world_theme", "fantasy")

            print(f"\n[bold green]Character will be linked to world: [/bold green][cyan]{linked_world.get('name')}[/cyan]")
            print(f"[bold green]Universe: [/bold green][cyan]{universe}[/cyan]")
            print(f"[bold green]World Theme: [/bold green][cyan]{theme}[/cyan]")

    # If not linking with a world or no world was selected, ask for universe and theme
    if not linked_world:
        universe = ask_with_examples(
            "[bold green]Choose your TTRPG universe[/bold green]",
            examples=["D&D", "Call of Cthulhu", "Warhammer 40K", "Shadowrun", "Star Wars RPG", "Vampire The Masquerade",
                      "Other"],
            default="D&D"
        )

        theme = ask_with_examples(
            "[bold green]Choose the universe theme for your story[/bold green]",
            examples=["Cyberpunk", "Gothic Horror", "Solarpunk", "Apocalypse", "Noir", "Urban Fantasy"],
            default=None
        )

    name = Prompt.ask("[bold green]What's your character's name?[/bold green]", default="Captain Kirk")

    gender = Prompt.ask("[bold green]What's your character's gender?[/bold green]", choices=["male", "female", "other"], default="male")

    race = ask_with_examples(
        "[bold green]Choose your character's race[/bold green]",
        examples=["Human", "Elf", "Dwarf", "Orc", "Other"],
        default="Human"
    )

    personality = ask_with_examples(
        "[bold green]Describe your character's personality (Optional)[/bold green]",
        examples=["Impulsive", "Passionate", "Quiet", "Overly Analytical"],
        default=None
    )

    appearance = ask_with_examples(
        "[bold green]Describe your character's appearance (Optional)[/bold green]",
        examples=["Short black hair, with big cheeks and long neck"],
        default=None
    )

    tone = ask_with_examples(
        "[bold green]Choose the tone for your story[/bold green]",
        examples=["Happy", "Sad", "Dark", "Hopeful", "Mysterious", "Epic"],
        default="Epic"
    )

    has_custom_story = Confirm.ask(
        "[bold yellow]Do you have a story you'd like to enhance?[/bold yellow]",
        default=False
    )

    custom_story = multiline_input(
        "Enter your character's story (press Enter twice to finish)") if has_custom_story else None

    return {
        "name": name,
        "gender": gender,
        "race": race,
        "personality": personality,
        "appearance": appearance,
        "universe": universe,
        "world_theme": theme,
        "tone": tone,
        "custom_story": custom_story,
        "linked_world_id": linked_world.get("id") if linked_world else None
    }
