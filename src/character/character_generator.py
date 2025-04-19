from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.character.character_vector_store import CharacterVectorStore
from src.character.llm import LLMFactory
from src.character.prompts import get_universe, get_world_theme, get_tone_context, create_appearance_prompt, \
    create_personality_prompt, create_backstory_prompt


class CharacterGenerator:
    def __init__(self):
        self.llm = LLMFactory.create()
        self.parser = StrOutputParser()
        self.vector_db = CharacterVectorStore()


    def __execute_chain(self, info: dict[str, str]) -> None:
        chains = {
            'appearance': create_appearance_prompt(info['appearance']),
            'personality': create_personality_prompt(info['personality']),
            'backstory': create_backstory_prompt(info['custom_story']),
        }

        for key, chain in chains.items():
            info[key] = (chain | self.llm | self.parser).invoke(info)


    def generate(self, info: dict[str, str]) -> dict[str, str]:
        info['id'] = str(uuid4())
        info['universe'] = get_universe().get(info['universe'], info['universe'])
        info['world_theme'] = get_world_theme().get(info['world_theme'], info['world_theme'])
        info['tone'] = get_tone_context().get(info['tone'], info['tone'])

        self.__execute_chain(info)
        self.vector_db.store(info)

        return info
