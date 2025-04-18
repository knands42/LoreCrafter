import re

from rich.prompt import Prompt, Confirm


def clean_text(value: str) -> str:
    # Lowercase and remove all characters except a-z, 0-9, and spaces
    return re.sub(r'[^a-z0-9 ]+', '', value)


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


def ask_with_examples(prompt_text: str, examples: list[str], default: str = None) -> str:
    example_str = f"[dim]e.g., {', '.join(examples)}[/dim]"
    full_prompt = f"{prompt_text} {example_str}"
    return Prompt.ask(full_prompt, default=default)


def get_character_info():
    print("\n[bold cyan]Let's create your character![/bold cyan]")

    universe = ask_with_examples(
        "[bold green]Choose your TTRPG universe[/bold green]",
        examples=["D&D", "Call of Cthulhu", "Starfinder", "Other"],
        default="D&D"
    )

    name = Prompt.ask("[bold green]What's your character's name?[/bold green]")

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

    theme = ask_with_examples(
        "[bold green]Choose the universe theme for your story[/bold green]",
        examples=["Cyberpunk", "Gothic Horror", "Solarpunk", "Apocalypse", "Noir", "Urban Fantasy"],
        default=None
    )

    if theme is None:
        theme = "default"

    has_custom_story = Confirm.ask(
        "[bold yellow]Do you have a story you'd like to enhance?[/bold yellow]",
        default=False
    )

    custom_story = multiline_input(
        "Enter your character's story (press Enter twice to finish)") if has_custom_story else None

    return {
        "name": name,
        "race": race,
        "personality": personality,
        "appearance": appearance,
        "universe": clean_text(universe).lower(),
        "universe_theme": clean_text(theme).lower(),
        "tone": clean_text(tone).lower(),
        "custom_story": custom_story
    }
