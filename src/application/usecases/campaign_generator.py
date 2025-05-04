from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.adapter.output.llm import LLMFactory
from src.adapter.output.repository import WorldVectorStore
from src.application.prompts import create_campaign_setting_prompt, create_hidden_elements_prompt


class CampaignGenerator:
    def __init__(
        self,
        world_vector_db: WorldVectorStore,
    ):
        self.llm = LLMFactory.create_chat()
        self.parser = StrOutputParser()
        self.world_vector_db = world_vector_db

    def generate(self, info: dict[str, str]) -> dict[str, str]:
        """Generate campaign and hidden elements for a world.

        Args:
            world_info: A dictionary containing world information.

        Returns:
            A dictionary containing campaign and hidden elements.
        """
        campaign_info = {key: value for key, value in info.items()}
        
        campaign_info['id'] = str(uuid4())

        self.__execute_chain(campaign_info)
        
        return campaign_info

    def __execute_chain(self, campaign_info: dict[str, str]) -> None:
        """Execute the chain to generate campaign and hidden elements.

        Args:
            campaign_info: A dictionary containing campaign information.
        """

        linked_world = campaign_info.get("linked_world", {})
        if isinstance(linked_world, dict):
            flat_campaign_info = {**campaign_info, **linked_world}
            flat_campaign_info.pop("linked_world", None)
        else:
            flat_campaign_info = campaign_info
            
        print(flat_campaign_info)

        campaign_chain = create_campaign_setting_prompt(campaign_info.get('custom_campaign'))
        campaign_info['campaign'] = (campaign_chain | self.llm | self.parser).invoke(flat_campaign_info)

        hidden_elements_chain = create_hidden_elements_prompt()
        campaign_info['hidden_elements'] = (hidden_elements_chain | self.llm | self.parser).invoke(campaign_info)