import os
from typing import Optional

from dotenv import load_dotenv
from langchain.schema.output_parser import StrOutputParser
from langchain_openai import ChatOpenAI

from .prompts import get_universe, get_universe_theme, get_tone_context, create_enhancement_prompt, \
    create_creation_prompt

load_dotenv()


def create_llm() -> ChatOpenAI:
    """Create and return a ChatOpenAI instance."""
    return ChatOpenAI(
        model="gpt-3.5-turbo",
        temperature=0.8,
        api_key=os.getenv("OPENAI_API_KEY")
    )


def generate_character(character_info: dict[str, str], custom_story: Optional[str] = None) -> str:
    """Generate or enhance a character's story based on the provided information."""
    universe = get_universe().get(character_info['universe'], character_info['universe'])
    theme = get_universe_theme().get(character_info['universe_theme'], character_info['universe_theme'])
    tone = get_tone_context().get(character_info['tone'], character_info['tone'])

    context = {
        "name": character_info["name"],
        "race": character_info["race"],
        "personality": character_info["personality"],
        "universe": universe,
        "universe_theme": theme,
        "tone": tone,
        "custom_story": custom_story,
    }

    llm = create_llm()
    prompt = create_enhancement_prompt() if custom_story else create_creation_prompt()

    chain = prompt | llm | StrOutputParser()
    result = chain.invoke(context)

    return result
