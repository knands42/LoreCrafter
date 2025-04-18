import os

from dotenv import load_dotenv
from langchain.schema.output_parser import StrOutputParser
from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_openai import ChatOpenAI

from .prompts import get_universe, get_world_theme, get_tone_context, create_backstory_prompt, \
    create_personality_prompt, create_appearance_prompt

load_dotenv()


def create_llm() -> ChatGoogleGenerativeAI | ChatOpenAI:
    """Create and return a llm instance."""
    if os.getenv("GOOGLE_API_KEY") is not None:
        return ChatGoogleGenerativeAI(
            model="gemini-2.0-flash",
            temperature=0.8,
            api_key=os.getenv("GOOGLE_API_KEY")
        )

    return ChatOpenAI(
        model="gpt-3.5-turbo",
        temperature=0.8,
        api_key=os.getenv("OPENAI_API_KEY")
    )


def execute_chain(character_info: dict[str, str], context: dict[str, str]) -> None:
    llm = create_llm()
    parser = StrOutputParser()

    appearance_chain = create_appearance_prompt(character_info['appearance']) | llm | parser
    personality_chain = create_personality_prompt(character_info['personality']) | llm | parser
    backstory_chain = create_backstory_prompt(character_info['custom_story']) | llm | parser

    context['appearance'] = appearance_chain.invoke(context)
    context['personality'] = personality_chain.invoke(context)
    context['backstory'] = backstory_chain.invoke(context)


def generate_character(character_info: dict[str, str]) -> dict[str, str]:
    """Generate or enhance a character's story based on the provided information."""
    universe = get_universe().get(character_info['universe'], character_info['universe'])
    theme = get_world_theme().get(character_info['world_theme'], character_info['world_theme'])
    tone = get_tone_context().get(character_info['tone'], character_info['tone'])

    context = {
        "name": character_info["name"],
        "race": character_info["race"],
        "universe": universe,
        "world_theme": theme,
        "tone": tone,
        "custom_story": character_info['custom_story'],
    }

    execute_chain(character_info, context)
    return context
