from uuid import uuid4

from langchain_core.output_parsers import StrOutputParser

from src.adapter.output.llm import LLMFactory
from src.adapter.output.repository import WorldVectorStore
from src.application.domain.campaign_domain import CampaignCreationDomain, CampaignDomain
from src.application.prompts import create_campaign_setting_prompt, create_hidden_elements_prompt


class CampaignGenerator:
    def __init__(
        self,
        world_vector_db: WorldVectorStore,
    ):
        self.llm = LLMFactory.create_chat()
        self.parser = StrOutputParser()
        self.world_vector_db = world_vector_db

    def generate(self, campaign_creation_domain: CampaignCreationDomain) -> CampaignDomain:
        """Generate campaign and hidden elements for a world.

        Args:
            campaign_creation_domain: A dictionary containing campaign information.

        Returns:
            A dictionary containing campaign and hidden elements.
        """

        generated_campaign, generated_hidden_elements = self.__generate_campaign(campaign_creation_domain)

        campaign_domain = CampaignDomain(
            id=uuid4(),
            name=campaign_creation_domain['name'],
            campaign=generated_campaign,
            hidden_elements=generated_hidden_elements,
            universe=campaign_creation_domain['universe'],
            world_theme=campaign_creation_domain['world_theme'],
            tone=campaign_creation_domain['tone']
        )

        return campaign_domain

    def __generate_campaign(self, campaign_creation_domain: CampaignCreationDomain):
        linked_world = campaign_creation_domain.get("linked_world", {})
        if isinstance(linked_world, dict):
            flat_campaign_info = {**campaign_creation_domain, **linked_world}
            flat_campaign_info.pop("linked_world", None)
        else:
            flat_campaign_info = campaign_creation_domain

        campaign_chain = create_campaign_setting_prompt(campaign_creation_domain.get('campaign'))
        generated_campaign = (campaign_chain | self.llm | self.parser).invoke(flat_campaign_info)

        hidden_elements_chain = create_hidden_elements_prompt()
        generated_hidden_elements = (hidden_elements_chain | self.llm | self.parser).invoke(flat_campaign_info)

        return generated_campaign, generated_hidden_elements
