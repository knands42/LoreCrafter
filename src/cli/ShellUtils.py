import re
from abc import ABC

from rich.prompt import Prompt


class ShellUtils(ABC):
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
