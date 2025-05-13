import base64
from pathlib import Path
from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.adapter.output.llm.llm_factory import LLMFactory
from src.adapter.output.vector_db.world_vector_store import WorldVectorStore
from src.application.domain.word_domain import World, WorldCreation
from src.application.prompts import get_world_theme, get_universe, get_story_tone, create_world_history_prompt, \
    create_timeline_prompt, get_world_image_prompt


class WorldGenerator:
    def __init__(
        self,
        vector_db: WorldVectorStore,
    ):
        self.llm = LLMFactory.create_chat()
        self.image_llm = LLMFactory.create_image_generator()
        self.parser = StrOutputParser()
        self.vector_db = vector_db

    def generate(self, world_create_domain: WorldCreation) -> World:
        world_create_domain.world_theme = get_world_theme(world_create_domain.world_theme)
        world_create_domain.universe = get_universe(world_create_domain.universe)
        world_create_domain.tone = get_story_tone(world_create_domain.tone)

        history_chain = create_world_history_prompt(world_create_domain.backstory)
        created_history = (history_chain | self.llm | self.parser).invoke(world_create_domain.__dict__)

        timeline_chain = create_timeline_prompt(world_create_domain.timeline)
        timeline_history = (timeline_chain | self.llm | self.parser).invoke(world_create_domain.__dict__)

        image_filename = self.__generate_image(created_history)

        world_domain = World(
            id=uuid4(),
            name=world_create_domain.name,
            universe=world_create_domain.universe,
            world_theme=world_create_domain.world_theme,
            tone=world_create_domain.tone,
            backstory=created_history,
            timeline=timeline_history,
            image_filename=image_filename
        )
        self.vector_db.store(world_domain)

        return world_domain

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
        image_filename = f"world_image.png"
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
