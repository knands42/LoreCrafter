import base64
from pathlib import Path
from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.adapter.output.llm import LLMFactory
from src.adapter.output.repository import CharacterVectorStore
from src.application.domain.character_domain import CharacterCreateDomain, CharacterDomain
from src.application.prompts import get_world_theme, get_universe, get_story_tone, create_appearance_prompt, \
    get_character_image_prompt


class CharacterGenerator:
    def __init__(
        self,
        vector_db: CharacterVectorStore,
    ):
        self.llm = LLMFactory.create_chat()
        self.image_llm = LLMFactory.create_image_generator()
        self.parser = StrOutputParser()
        self.vector_db = vector_db

    def generate(self, character_creation_info: CharacterCreateDomain) -> CharacterDomain:
        character_creation_info['world_theme'] = get_world_theme(character_creation_info['world_theme'])
        character_creation_info['universe'] = get_universe(character_creation_info['universe'])
        character_creation_info['tone'] = get_story_tone(character_creation_info['tone'])

        appearance_chain = create_appearance_prompt(character_creation_info['appearance'])
        generated_appearance = (appearance_chain | self.llm | self.parser).invoke(character_creation_info)

        personality_chain = create_appearance_prompt(character_creation_info['personality'])
        generated_personality = (personality_chain | self.llm | self.parser).invoke(character_creation_info)

        backstory_chain = create_appearance_prompt(character_creation_info['backstory'])
        generated_backstory = (backstory_chain | self.llm | self.parser).invoke(character_creation_info)
        image_filename = self.__generate_image(generated_appearance)

        character_domain = CharacterDomain(
            id=uuid4(),
            name=character_creation_info['name'],
            race=character_creation_info['race'],
            gender=character_creation_info['gender'],

            appearance=generated_appearance,
            personality=generated_personality,
            backstory=generated_backstory,

            world_theme=character_creation_info['world_theme'],
            universe=character_creation_info['universe'],
            tone=character_creation_info['tone'],

            linked_world_id=character_creation_info['linked_world_id'],
            image_filename=image_filename,
        )

        self.vector_db.store(character_domain)

        return character_domain

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
