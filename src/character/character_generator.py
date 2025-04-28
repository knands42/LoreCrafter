from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.character.character_vector_store import CharacterVectorStore
from src.character.llm import LLMFactory
from src.character.prompts import create_appearance_prompt, create_personality_prompt, create_backstory_prompt, \
    get_world_theme, get_story_tone, get_universe


class CharacterGenerator:
    def __init__(self):
        self.llm = LLMFactory.create()
        self.parser = StrOutputParser()
        self.vector_db = CharacterVectorStore()

    def generate(self, info: dict[str, str]) -> dict[str, str]:
        character_info = {key: value for key, value in info.items()}

        character_info['id'] = str(uuid4())
        character_info['world_theme_prompt'] = get_world_theme(character_info['world_theme'])
        character_info['universe_prompt'] = get_universe(character_info['universe'])
        character_info['tone_prompt'] = get_story_tone(character_info['tone'])

        self.__execute_chain(character_info)
        self.vector_db.store(character_info)
        self.__generate_image(character_info.get("appearance"))

        return character_info

    def __execute_chain(self, character_info: dict[str, str]) -> None:
        chains = {
            'appearance': create_appearance_prompt(character_info['appearance']),
            'personality': create_personality_prompt(character_info['personality']),
            'backstory': create_backstory_prompt(character_info['custom_story']),
        }

        for key, chain in chains.items():
            character_info[key] = (chain | self.llm | self.parser).invoke(character_info)

    def __generate_image(self, appearance_description: str):
        pass
