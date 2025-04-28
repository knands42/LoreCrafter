from weasyprint import HTML
from rich import print
from src.content_generator.template_manager import TemplateManager

# Create an instance of the TemplateManager
template_manager = TemplateManager()

# Main template function
def get_template(character_info: dict[str, any]) -> str:
    """Get the appropriate template based on character info."""
    return template_manager.get_template(character_info)

# Create the PDF
def create_pdf(content: dict):
    print("\n[bold yellow]Generating your character sheet...[/bold yellow]")
    html_content = get_template(content)
    HTML(string=html_content).write_pdf("assets/character_profile.pdf", presentational_hints=True)
    print("\n[bold yellow]Character sheet ready[/bold yellow]")