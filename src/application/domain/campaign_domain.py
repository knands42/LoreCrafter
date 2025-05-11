from typing import Optional, List
from uuid import UUID

from pydantic import BaseModel, Field

from src.application.domain.character_domain import Character
from src.application.domain.word_domain import World


class CampaignCreation(BaseModel):
    """Schema for creating a new campaign"""
    name: str = Field(..., description="The campaign's name")
    universe: str = Field(..., description="The universe the campaign belongs to")
    world_theme: str = Field(..., description="The theme of the campaign's world")
    tone: str = Field(..., description="The tone of the campaign")
    campaign: Optional[str] = Field(None, description="The campaign's description")
    linked_world: Optional[World] = Field(None, description="A linked world, if any")
    linked_character: List[Character] = Field([], description="List of linked characters")

    class Config:
        json_schema_extra = {
            "example": {
                "name": "The Crystal Prophecy",
                "universe": "Fantasy",
                "world_theme": "High Fantasy",
                "tone": "Epic",
                "campaign": "A quest to recover the shattered pieces of the Crystal of Eternity...",
                "linked_world": None,
                "linked_character": []
            }
        }


class Campaign(BaseModel):
    """Schema for a campaign with all attributes"""
    id: UUID = Field(..., description="Unique identifier for the campaign")
    name: str = Field(..., description="The campaign's name")
    universe: str = Field(..., description="The universe the campaign belongs to")
    world_theme: str = Field(..., description="The theme of the campaign's world")
    tone: str = Field(..., description="The tone of the campaign")
    campaign: Optional[str] = Field(None, description="The campaign's description")
    hidden_elements: Optional[str] = Field(None, description="Hidden elements of the campaign")

    class Config:
        json_schema_extra = {
            "example": {
                "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
                "name": "The Crystal Prophecy",
                "universe": "Fantasy",
                "world_theme": "High Fantasy",
                "tone": "Epic",
                "campaign": "A quest to recover the shattered pieces of the Crystal of Eternity...",
                "hidden_elements": "The true villain is actually the king's advisor..."
            }
        }
