from rich.panel import Panel
from rich.text import Text
from rich import print


def print_character(character: dict[str, str | None]):
    def format_section(title: str, content: str) -> Panel:
        return Panel.fit(
            Text(content.strip(), style="italic"),
            title=title,
            border_style="cyan",
            padding=(1, 2),
            title_align="left"
        )

    print("[bold magenta]✨ Character Overview ✨[/bold magenta]\n")

    print(f"[bold]Name:[/bold] {character.get('name') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Race:[/bold] {character.get('race', '[dim]Unknown[/dim]')}")
    print(f"[bold]Universe:[/bold] {character.get('universe') or '[dim]Unknown[/dim]'}")
    print(f"[bold]Theme:[/bold] {character.get('world_theme', '[dim]Unknown[/dim]')}")
    print(f"[bold]Tone:[/bold] {character.get('tone', '[dim]Unknown[/dim]')}")
    print()

    if character.get("appearance"):
        print(format_section("Appearance", character["appearance"]))

    if character.get("personality"):
        print(format_section("Personality", character["personality"]))

    if character.get("backstory"):
        print(format_section("Backstory", character["backstory"]))
