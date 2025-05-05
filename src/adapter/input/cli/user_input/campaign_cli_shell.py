from typing import List

from rich import print
from rich.console import Console
from rich.prompt import Prompt, Confirm

from src.adapter.input.cli.ShellUtils import ShellUtils
from src.adapter.output.repository.world_vector_store import WorldVectorStore
from src.application.domain.campaign_domain import CampaignCreationDomain


class CampaignCLIShell(ShellUtils):
    def __init__(self, console: Console, world_vector_store: WorldVectorStore):
        super().__init__(console, world_vector_store)

    def get_campaign_info(self, get_default: bool = False) -> CampaignCreationDomain:
        if get_default:
            return CampaignCreationDomain(
                name="The Lost Mines",
                universe="D&D",
                world_theme="fantasy",
                tone="Epic",
                campaign=None,
                linked_world=None,
                linked_character=[],
            )

        print("\n[bold magenta]Let's create your campaign![/bold magenta]")

        # Ask if the user wants to create a campaign in an existing world
        should_link_with_world = Confirm.ask(
            "[bold yellow]Do you want to create a campaign in an existing world?[/bold yellow]",
            default=True
        )

        universe = None
        theme = None
        selected_world = None

        if should_link_with_world:
            worlds = self.display_worlds()
            selected_world = self.select_world_by_id(worlds)

            if selected_world:
                # Use universe and world_theme from the linked world
                universe = selected_world.get("universe", "D&D")
                theme = selected_world.get("world_theme", "fantasy")

                print(
                    f"\n[bold green]Campaign will be created in world: [/bold green][cyan]{selected_world.get('name')}[/cyan]")
                print(f"[bold green]Universe: [/bold green][cyan]{universe}[/cyan]")
                print(f"[bold green]World Theme: [/bold green][cyan]{theme}[/cyan]")

        # If not linking with a world or no world was selected, ask for universe and theme
        if not selected_world:
            universe = self.ask_with_examples(
                "[bold green]Choose your TTRPG universe[/bold green]",
                examples=["D&D", "Call of Cthulhu", "Warhammer 40K", "Shadowrun", "Star Wars RPG",
                          "Vampire The Masquerade",
                          "Other"],
                default="D&D"
            )

            theme = self.ask_with_examples(
                "[bold green]Choose the universe theme for your campaign[/bold green]",
                examples=["Cyberpunk", "Gothic Horror", "Solarpunk", "Apocalypse", "Noir", "Urban Fantasy"],
                default="Fantasy"
            )

        name = Prompt.ask("[bold green]What's your campaign's name?[/bold green]", default="The Lost Mines")

        tone = self.ask_with_examples(
            "[bold green]Choose the tone for your campaign[/bold green]",
            examples=["Happy", "Sad", "Dark", "Hopeful", "Mysterious", "Epic"],
            default="Epic"
        )

        # Ask if the user wants to link characters to the campaign
        should_link_characters = Confirm.ask(
            "[bold yellow]Do you want to link existing characters to your campaign?[/bold yellow]",
            default=False
        )

        linked_character = []
        if should_link_characters:
            linked_character = self.__select_characters()

        has_custom_campaign = Confirm.ask(
            "[bold yellow]Do you have a campaign setting you'd like to enhance?[/bold yellow]",
            default=False
        )

        custom_campaign = self.multiline_input(
            "Enter your campaign setting (press Enter twice to finish)") if has_custom_campaign else None

        return CampaignCreationDomain(
            name=name,
            universe=universe.lower(),
            world_theme=theme.lower(),
            tone=tone.lower(),
            campaign=custom_campaign,
            linked_world=selected_world,
            linked_character=linked_character
        )

    def __select_characters(self) -> List[str]:
        """Display a list of available characters and let the user select multiple.

        Returns:
            A list of character IDs that were selected.
        """
        # This is a simplified version - in a real implementation, you would
        # query the character vector store to get all characters
        characters = self.__display_characters()

        if not characters:
            print("[yellow]No characters found. You'll create a campaign without linked characters.[/yellow]")
            return []

        selected_ids = []
        while True:
            selection = Prompt.ask(
                "[bold green]Select a character by number (or press Enter to finish)[/bold green]",
                default=None
            )

            if not selection:
                break

            try:
                index = int(selection) - 1
                if 0 <= index < len(characters):
                    char_id = characters[index].get("id")
                    if char_id and char_id not in selected_ids:
                        selected_ids.append(char_id)
                        print(f"[green]Added {characters[index].get('name')} to the campaign[/green]")
                    else:
                        print("[yellow]Character already added or has no ID[/yellow]")
                else:
                    print("[red]Invalid selection.[/red]")
            except ValueError:
                print("[red]Invalid input.[/red]")

        return selected_ids

    def __display_characters(self) -> list[dict]:
        """Display a list of available characters.
        
        Returns:
            A list of character dictionaries.
        """
        # In a real implementation, you would query the character vector store
        # This is a placeholder that would need to be implemented
        # For now, we'll return an empty list
        print("[yellow]Character selection is not fully implemented yet.[/yellow]")
        return []
