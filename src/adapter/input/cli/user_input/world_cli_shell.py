from rich import print
from rich.console import Console
from rich.prompt import Prompt, Confirm

from src.adapter.input.cli.ShellUtils import ShellUtils
from src.adapter.output.repository import WorldVectorStore
from src.application.domain.word_domain import WorldCreateDomain


class WorldCLIShell(ShellUtils):
    def __init__(self, console: Console, vector_store: WorldVectorStore):
        super().__init__(console, vector_store)
        self.console = Console()

    def get_world_info(self, get_default: bool = False) -> WorldCreateDomain:
        if get_default:
            return {
                "name": "Eldoria",
                "universe": "D&D",
                "world_theme": "fantasy",
                "tone": "Epic",
                "backstory": None,
                "timeline": None,
            }

        print("\n[bold magenta]Let's create your world![/bold magenta]")

        universe = self.ask_with_examples(
            "[bold green]Choose your TTRPG universe[/bold green]",
            examples=["D&D", "Call of Cthulhu", "Warhammer 40K", "Shadowrun", "Star Wars RPG", "Vampire The Masquerade",
                      "Other"],
            default="D&D"
        )

        name = Prompt.ask("[bold green]What's your world's name?[/bold green]", default="Eldoria")

        tone = self.ask_with_examples(
            "[bold green]Choose the tone for your world[/bold green]",
            examples=["Happy", "Sad", "Dark", "Hopeful", "Mysterious", "Epic"],
            default="Epic"
        )

        theme = self.ask_with_examples(
            "[bold green]Choose the theme for your world[/bold green]",
            examples=["Cyberpunk", "Gothic Horror", "Solarpunk", "Apocalypse", "Noir", "Urban Fantasy", "Fantasy"],
            default="Fantasy"
        )

        has_custom_history = Confirm.ask(
            "[bold yellow]Do you have a world history you'd like to enhance?[/bold yellow]",
            default=False
        )

        custom_history = self.multiline_input(
            "Enter your world's history (press Enter twice to finish)") if has_custom_history else None

        has_custom_timeline = Confirm.ask(
            "[bold yellow]Do you have a timeline you'd like to enhance?[/bold yellow]",
            default=False
        )

        custom_timeline = self.multiline_input(
            "Enter your world's timeline (press Enter twice to finish)") if has_custom_timeline else None

        return {
            "name": name,
            "universe": universe,
            "world_theme": theme,
            "tone": tone,
            "backstory": custom_history,
            "timeline": custom_timeline,
        }
