from weasyprint import HTML
from rich import print

from src.adapter.output.pdf.template_manager import TemplateManager
from src.application.domain.character_domain import Character
from src.application.domain.word_domain import World

# Create an instance of the TemplateManager
template_manager = TemplateManager()

# Main template function
def get_template(character_info: Character) -> str:
    """Get the appropriate template based on character info."""
    return template_manager.get_template(character_info)

# Get world lore template
def get_world_template(world_info: World) -> str:
    """Get the appropriate template for world lore."""
    return template_manager.get_world_template(world_info)

# Create the character PDF
def create_character_pdf(content: Character, filename: str = "character_profile.pdf"):
    print("\n[bold yellow]Generating your character sheet...[/bold yellow]")
    html_content = get_template(content)
    HTML(string=html_content).write_pdf(f"assets/{filename}", presentational_hints=True)
    print(f"\n[bold yellow]Character sheet ready: assets/{filename}[/bold yellow]")

# Create the world lore PDF
def create_world_pdf(content: World, filename: str = "world_lore.pdf"):
    print("\n[bold yellow]Generating your world lore document...[/bold yellow]")
    html_content = get_world_template(content)
    HTML(string=html_content).write_pdf(f"assets/{filename}", presentational_hints=True)
    print(f"\n[bold yellow]World lore document ready: assets/{filename}[/bold yellow]")
