import uuid
from typing import TypedDict

from src.application.domain.character_domain import CharacterDomain
from src.application.domain.word_domain import WorldDomain


class CampaignDomain(TypedDict):
    id: uuid.UUID
    name: str
    universe: str
    world_theme: str
    tone: str
    campaign: str | None
    hidden_elements: str | None



class CampaignCreationDomain(TypedDict):
    name: str
    universe: str
    world_theme: str
    tone: str
    campaign: str | None
    linked_world: WorldDomain | None
    linked_character: list[CharacterDomain]
