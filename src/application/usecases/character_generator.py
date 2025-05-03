import base64
from pathlib import Path
from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.adapter.output.llm import LLMFactory
from src.adapter.output.repository import CharacterVectorStore
from src.application.prompts import get_world_theme, get_universe, get_story_tone, create_appearance_prompt, \
    create_personality_prompt, create_backstory_prompt, get_character_image_prompt


class CharacterGenerator:
    def __init__(
        self,
        vector_db: CharacterVectorStore,
    ):
        self.llm = LLMFactory.create(),
        self.image_llm = LLMFactory.create_image_generator()
        self.parser = StrOutputParser()
        self.vector_db = vector_db

    def generate(self, info: dict[str, str]) -> dict[str, str]:
        character_info = {key: value for key, value in info.items()}

        character_info['id'] = str(uuid4())
        character_info['world_theme_prompt'] = get_world_theme(character_info['world_theme'])
        character_info['universe_prompt'] = get_universe(character_info['universe'])
        character_info['tone_prompt'] = get_story_tone(character_info['tone'])

        self.__execute_chain(character_info)
        self.vector_db.store(character_info)

        # Generate character image and store the path
        image_filename = self.__generate_image(character_info.get("appearance"))
        if image_filename:
            character_info['image_filename'] = image_filename

        return character_info

    def __execute_chain(self, character_info: dict[str, str]) -> None:
        chains = {
            'appearance': create_appearance_prompt(character_info['appearance']),
            'personality': create_personality_prompt(character_info['personality']),
            'backstory': create_backstory_prompt(character_info['custom_story']),
        }

        for key, chain in chains.items():
            character_info[key] = (chain | self.llm | self.parser).invoke(character_info)

    def __generate_image(self, appearance_description: str) -> str | None:
        """Generate an image based on the character's appearance description and save it to the assets folder.

        Args:
            appearance_description: A detailed description of the character's appearance.

        Returns:
            The path to the generated image file.
        """
        if not appearance_description:
            return None

        assets_dir = Path("assets")
        assets_dir.mkdir(exist_ok=True)

        # Generate a unique filename for the image
        image_filename = f"character_image_{uuid4()}.png"
        image_path = assets_dir / image_filename

        try:
            prompt = get_character_image_prompt(appearance_description)
            response = self.image_llm.invoke(
                [prompt],
                generation_config=dict(response_modalities=["TEXT", "IMAGE"]),
            )

            image_base64 = response.content[0].get("image_url").get("url").split(",")[-1]
            if image_base64:
                with open(image_path, 'wb') as f:
                    f.write(base64.b64decode(image_base64))
                    return str(image_filename)

            print("No image data received from the API.")
            return None

        except Exception as e:
            print(f"Error generating image: {e}")
            return None
