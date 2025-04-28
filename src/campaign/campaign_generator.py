from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.character.llm import LLMFactory
from src.campaign.prompts import (
    create_campaign_setting_prompt,
    create_hidden_elements_prompt
)


class CampaignGenerator:
    def __init__(self):
        self.llm = LLMFactory.create()
        self.parser = StrOutputParser()

    def generate(self, world_info: dict[str, str]) -> dict[str, str]:
        """Generate campaign and hidden elements for a world.

        Args:
            world_info: A dictionary containing world information.

        Returns:
            A dictionary containing campaign and hidden elements.
        """
        campaign_info = {key: value for key, value in world_info.items()}
        
        world_info['id'] = str(uuid4())

        self.__execute_chain(campaign_info, world_info.get('custom_campaign'))
        
        return campaign_info

    def __execute_chain(self, campaign_info: dict[str, str], custom_campaign: str = None) -> None:
        """Execute the chain to generate campaign and hidden elements.

        Args:
            campaign_info: A dictionary containing campaign information.
            custom_campaign: A custom campaign setting to enhance.
        """
        campaign_chain = create_campaign_setting_prompt(custom_campaign)
        campaign_info['campaign'] = (campaign_chain | self.llm | self.parser).invoke(campaign_info)

        hidden_elements_chain = create_hidden_elements_prompt()
        campaign_info['hidden_elements'] = (hidden_elements_chain | self.llm | self.parser).invoke(campaign_info)