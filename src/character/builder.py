import os
from typing import Optional, Any

from dotenv import load_dotenv
from langchain.schema.output_parser import StrOutputParser
from langchain.schema.runnable import RunnableMap, RunnableSequence, Runnable
from langchain_core.prompts import PromptTemplate
from langchain_core.runnables import RunnableConfig, RunnableLambda
from langchain_core.runnables.utils import Input, Output
from langchain_openai import ChatOpenAI

from .prompts import get_universe, get_universe_theme, get_tone_context, create_backstory_prompt, \
    create_personality_prompt, create_appearance_prompt

load_dotenv()


def create_llm() -> ChatOpenAI:
    """Create and return a ChatOpenAI instance."""
    return ChatOpenAI(
        model="gpt-3.5-turbo",
        temperature=0.8,
        api_key=os.getenv("OPENAI_API_KEY")
    )


def execute_chain(character_info: dict[str, str], context: dict[str, str]) -> str:
    llm = create_llm()
    parser = StrOutputParser()

    appearance_chain = create_appearance_prompt(character_info['appearance']) | llm | parser
    personality_chain = create_personality_prompt(character_info['personality']) | llm | parser
    backstory_chain = create_backstory_prompt(character_info['custom_story']) | llm | parser

    appearance_result = appearance_chain.invoke(context)
    context['appearance'] = appearance_result
    personality_result = personality_chain.invoke(context)
    context['personality'] = personality_result

    return backstory_chain.invoke(context)
    

def generate_character(character_info: dict[str, str]) -> str:
    """Generate or enhance a character's story based on the provided information."""
    universe = get_universe().get(character_info['universe'], character_info['universe'])
    theme = get_universe_theme().get(character_info['universe_theme'], character_info['universe_theme'])
    tone = get_tone_context().get(character_info['tone'], character_info['tone'])

    context = {
        "name": character_info["name"],
        "race": character_info["race"],
        "universe": universe,
        "universe_theme": theme,
        "tone": tone,
        "custom_story": character_info['custom_story'],
    }

    return execute_chain(character_info, context)


