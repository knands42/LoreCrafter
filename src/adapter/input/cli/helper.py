from rich.panel import Panel
from rich.text import Text
from rich.markdown import Markdown
from rich import print

from src.application.domain.campaign_domain import Campaign
from src.application.domain.character_domain import Character
from src.application.domain.word_domain import World


def format_section(title: str, content: str, use_markdown: bool = False) -> Panel:
    """Format a section of text with a title and border.

    Args:
        title: The title of the section
        content: The content of the section
        use_markdown: Whether to render the content as Markdown

    Returns:
        A Panel containing the formatted content
    """
    content_obj = Markdown(content.strip()) if use_markdown else Text(content.strip(), style="italic")
    return Panel.fit(
        content_obj,
        title=title,
        border_style="cyan",
        padding=(1, 2),
        title_align="left"
    )


def print_character(character: Character):
    print("[bold magenta]✨ Character Overview ✨[/bold magenta]\n")

    print(f"[bold]Name:[/bold] {character.get('name') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Race:[/bold] {character.get('race') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Universe:[/bold] {character.get('universe') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Theme:[/bold] {character.get('world_theme') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Tone:[/bold] {character.get('tone') or '[dim]Unknown[/dim]'}")
    print()

    if character.get("appearance"):
        print(format_section("Appearance", character["appearance"]))

    if character.get("personality"):
        print(format_section("Personality", character["personality"]))

    if character.get("backstory"):
        print(format_section("Backstory", character["backstory"], use_markdown=True))


def print_world(world: World):
    print("[bold magenta]✨ World Overview ✨[/bold magenta]\n")

    print(f"[bold]Name:[/bold] {world.get('name') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Universe:[/bold] {world.get('universe') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Theme:[/bold] {world.get('world_theme') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Tone:[/bold] {world.get('tone') or '[dim]Unknown[/dim]'}")
    print()

    if world.get("backstory"):
        print(format_section("World History", world["backstory"], use_markdown=True))

    if world.get("timeline"):
        print(format_section("Timeline", world["timeline"], use_markdown=True))


def print_campaign(campaign: Campaign):
    print("[bold magenta]✨ Campaign Overview ✨[/bold magenta]\n")

    print(f"[bold]Name:[/bold] {campaign.get('name') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Universe:[/bold] {campaign.get('universe') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Theme:[/bold] {campaign.get('world_theme') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Tone:[/bold] {campaign.get('tone') or '[dim]Unknown[/dim]'}")
    print()

    if campaign.get("campaign"):
        print(format_section("Campaign Setting", campaign["campaign"], use_markdown=True))

    if campaign.get("hidden_elements"):
        print(format_section("Hidden Elements", campaign["hidden_elements"], use_markdown=True))
