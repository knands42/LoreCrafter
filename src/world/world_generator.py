from uuid import uuid4
import os
import base64
from pathlib import Path

from langchain_core.output_parsers import StrOutputParser

from src.character.llm import LLMFactory
from src.character.prompts import get_world_theme, get_story_tone, get_universe
from src.world.prompts import (
    create_world_history_prompt,
    create_timeline_prompt,
    create_campaign_setting_prompt,
    create_hidden_elements_prompt,
    get_world_image_prompt
)
from src.world.world_vector_store import WorldVectorStore


class WorldGenerator:
    def __init__(self):
        self.llm = LLMFactory.create()
        self.image_llm = LLMFactory.create_image_generator()
        self.parser = StrOutputParser()
        self.vector_db = WorldVectorStore()

    def generate(self, info: dict[str, str]) -> dict[str, str]:
        world_info = {key: value for key, value in info.items()}

        world_info['id'] = str(uuid4())
        world_info['world_theme_prompt'] = get_world_theme(world_info['world_theme'])
        world_info['universe_prompt'] = get_universe(world_info['universe'])
        world_info['tone_prompt'] = get_story_tone(world_info['tone'])

        self.__execute_chain(world_info)
        self.vector_db.store(world_info)

        # Generate world image and store the path
        image_filename = self.__generate_image(world_info.get("history"))
        if image_filename:
            world_info['image_filename'] = image_filename

        return world_info

    def __execute_chain(self, world_info: dict[str, str]) -> None:
        history_chain = create_world_history_prompt(world_info.get('custom_history'))
        world_info['history'] = (history_chain | self.llm | self.parser).invoke(world_info)

        timeline_chain = create_timeline_prompt(world_info.get('custom_timeline'))
        world_info['timeline'] = (timeline_chain | self.llm | self.parser).invoke(world_info)

        campaign_chain = create_campaign_setting_prompt(world_info.get('custom_setting'))
        world_info['campaign'] = (campaign_chain | self.llm | self.parser).invoke(world_info)

        hidden_elements_chain = create_hidden_elements_prompt()
        world_info['hidden_elements'] = (hidden_elements_chain | self.llm | self.parser).invoke(world_info)

    def __generate_image(self, world_description: str) -> str | None:
        """Generate an image based on the world description and save it to the assets folder.

        Args:
            world_description: A detailed description of the world.

        Returns:
            The path to the generated image file.
        """
        if not world_description:
            return None

        assets_dir = Path("assets")
        assets_dir.mkdir(exist_ok=True)

        # Generate a unique filename for the image
        image_filename = f"world_image_{uuid4()}.png"
        image_path = assets_dir / image_filename

        try:
            prompt = get_world_image_prompt(world_description)
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