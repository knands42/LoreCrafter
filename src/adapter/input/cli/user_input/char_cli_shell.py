from typing import Optional

from rich import print
from rich.console import Console
from rich.prompt import Prompt, Confirm

from src.adapter.input.cli.user_input.ShellUtils import ShellUtils
from src.adapter.output.repository.world_vector_store import WorldVectorStore
from src.application.domain.character_domain import CharacterCreation


class CharCLIShell(ShellUtils):
    def __init__(self, console: Console, vector_store: WorldVectorStore):
        super().__init__(console, vector_store)
        self.vector_store = vector_store
        self.console = Console()

    def get_character_info(self, get_default: bool = False) -> CharacterCreation:
        if get_default:
            return CharacterCreation(
                name="Captain Kirk",
                gender="male",
                race="Human",
                personality=None,
                appearance=None,
                universe="D&D",
                world_theme="fantasy",
                tone="Epic",
                backstory=None,
                linked_world_id=None
            )

        print("\n[bold magenta]Let's create your character![/bold magenta]")

        # Ask if the user wants to link with an existing world
        should_link_with_world = Confirm.ask(
            "[bold yellow]Do you want to link your character with an existing world?[/bold yellow]",
            default=False
        )

        linked_world = None
        universe = None
        theme = None
        worlds = None

        if should_link_with_world:
            worlds = self.display_worlds()
            linked_world = self.select_world_by_id(worlds)

            if linked_world:
                # Use universe and world_theme from the linked world
                universe = linked_world.get("universe", "D&D")
                theme = linked_world.get("world_theme", "fantasy")

                print(
                    f"\n[bold green]Character will be linked to world: [/bold green][cyan]{linked_world.get('name')}[/cyan]")
                print(f"[bold green]Universe: [/bold green][cyan]{universe}[/cyan]")
                print(f"[bold green]World Theme: [/bold green][cyan]{theme}[/cyan]")

        # If not linking with a world or no world was selected, ask for universe and theme
        if not linked_world or worlds is None:
            universe = self.ask_with_examples(
                "[bold green]Choose your TTRPG universe[/bold green]",
                examples=["D&D", "Call of Cthulhu", "Warhammer 40K", "Shadowrun", "Star Wars RPG",
                          "Vampire The Masquerade",
                          "Other"],
                default="D&D"
            )

            theme = self.ask_with_examples(
                "[bold green]Choose the universe theme for your story[/bold green]",
                examples=["Cyberpunk", "Gothic Horror", "Solarpunk", "Apocalypse", "Noir", "Urban Fantasy"],
                default=None
            )

        name = Prompt.ask("[bold green]What's your character's name?[/bold green]", default="Captain Kirk")

        gender = Prompt.ask("[bold green]What's your character's gender?[/bold green]",
                            choices=["male", "female", "other"],
                            default="male")

        race = self.ask_with_examples(
            "[bold green]Choose your character's race[/bold green]",
            examples=["Human", "Elf", "Dwarf", "Orc", "Other"],
            default="Human"
        )

        personality = self.ask_with_examples(
            "[bold green]Describe your character's personality (Optional)[/bold green]",
            examples=["Impulsive", "Passionate", "Quiet", "Overly Analytical"],
            default=None
        )

        appearance = self.ask_with_examples(
            "[bold green]Describe your character's appearance (Optional)[/bold green]",
            examples=["Short black hair, with big cheeks and long neck"],
            default=None
        )

        tone = self.ask_with_examples(
            "[bold green]Choose the tone for your story[/bold green]",
            examples=["Happy", "Sad", "Dark", "Hopeful", "Mysterious", "Epic"],
            default="Epic"
        )

        has_custom_story = Confirm.ask(
            "[bold yellow]Do you have a story you'd like to enhance?[/bold yellow]",
            default=False
        )

        custom_story = self.multiline_input(
            "Enter your character's story (press Enter twice to finish)") if has_custom_story else None

        return CharacterCreation(
            name=name,
            gender=gender,
            race=race,
            personality=personality,
            appearance=appearance,
            universe=universe if universe else "d&d",
            world_theme=theme if theme else "fantasy",
            tone=tone if tone else "epic",
            backstory=custom_story,
            linked_world_id=linked_world.get("id") if linked_world else None
        )
