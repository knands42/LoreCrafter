import re
from abc import ABC
from typing import Optional

from rich.console import Console
from rich.prompt import Prompt
from rich.table import Table

from src.adapter.output.vector_db.world_vector_store import WorldVectorStore


class ShellUtils(ABC):
    def __init__(self, console: Console, vector_store: WorldVectorStore):
        self.console = console
        self.vector_store = vector_store
        
    
    def clean_text(self, value: str) -> str:
        # remove all characters except a-z, 0-9, and spaces
        return re.sub(r'[^a-zA-Z0-9 ]+', '', value)

    def multiline_input(self, prompt: str) -> str:
        """Capture multiline input until the user presses Enter twice."""
        print(f"\n[bold cyan]{prompt}[/bold cyan]")
        lines = []
        while True:
            line = input()
            if not line and lines and not lines[-1]:
                break
            lines.append(line)
        return "\n".join(lines[:-1])  # Remove the last empty line

    def ask_with_examples(self, prompt_text: str, examples: list[str], default: str = None) -> str:
        example_str = f"[dim]e.g., {', '.join(examples)}[/dim]"
        full_prompt = f"{prompt_text} {example_str}"
        return Prompt.ask(full_prompt, default=default)

    def display_worlds(self) -> list[dict]:
        """Display a list of available worlds and let the user select one.

        Returns:
            The selected world as a dictionary, or None if no world is selected.
        """
        worlds = self.vector_store.get_all_worlds()

        if not worlds:
            print("[yellow]No worlds found. You'll need to create a character without linking to a world.[/yellow]")
            return []

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

        self.console.print(table)

        return worlds

    def select_world_by_id(self, worlds: list[dict] | None) -> Optional[dict]:
        if not worlds:
            print("[red]No existing worlds found. You'll need to create a campaign without linking to a world.[/red]")
            return None

        selection = Prompt.ask(
            "[bold green]Select a world by number (or press Enter to skip)[/bold green]",
            default=None
        )

        if not selection:
            return None

        try:
            index = int(selection) - 1
            if 0 <= index < len(worlds):
                return worlds[index]
            else:
                print("[red]Invalid selection. Creating campaign without linking to a world.[/red]")
                return None
        except ValueError:
            print("[red]Invalid input. Creating campaign without linking to a world.[/red]")
            return None