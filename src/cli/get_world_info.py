from rich import print
from rich.prompt import Prompt, Confirm

from src.cli.get_char_info import multiline_input, ask_with_examples


def get_world_info(get_default: bool = False):
    if get_default:
        return {
            "name": "Eldoria",
            "universe": "D&D",
            "world_theme": "fantasy",
            "tone": "Epic",
            "custom_history": None,
            "custom_timeline": None,
            "custom_campaign": None
        }

    print("\n[bold magenta]Let's create your world![/bold magenta]")

    universe = ask_with_examples(
        "[bold green]Choose your TTRPG universe[/bold green]",
        examples=["D&D", "Call of Cthulhu", "Warhammer 40K", "Shadowrun", "Star Wars RPG", "Vampire The Masquerade",
                  "Other"],
        default="D&D"
    )

    name = Prompt.ask("[bold green]What's your world's name?[/bold green]", default="Eldoria")

    tone = ask_with_examples(
        "[bold green]Choose the tone for your world[/bold green]",
        examples=["Happy", "Sad", "Dark", "Hopeful", "Mysterious", "Epic"],
        default="Epic"
    )

    theme = ask_with_examples(
        "[bold green]Choose the theme for your world[/bold green]",
        examples=["Cyberpunk", "Gothic Horror", "Solarpunk", "Apocalypse", "Noir", "Urban Fantasy", "Fantasy"],
        default="Fantasy"
    )

    has_custom_history = Confirm.ask(
        "[bold yellow]Do you have a world history you'd like to enhance?[/bold yellow]",
        default=False
    )

    custom_history = multiline_input(
        "Enter your world's history (press Enter twice to finish)") if has_custom_history else None

    has_custom_timeline = Confirm.ask(
        "[bold yellow]Do you have a timeline you'd like to enhance?[/bold yellow]",
        default=False
    )

    custom_timeline = multiline_input(
        "Enter your world's timeline (press Enter twice to finish)") if has_custom_timeline else None

    has_custom_setting = Confirm.ask(
        "[bold yellow]Do you have a campaign setting you'd like to enhance?[/bold yellow]",
        default=False
    )

    custom_setting = multiline_input(
        "Enter your campaign setting (press Enter twice to finish)") if has_custom_setting else None

    return {
        "name": name,
        "universe": universe,
        "world_theme": theme,
        "tone": tone,
        "custom_history": custom_history,
        "custom_timeline": custom_timeline,
        "custom_setting": custom_setting
    }
