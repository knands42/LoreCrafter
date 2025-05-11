from typing import Optional, List
from uuid import UUID

from pydantic import BaseModel, Field
from typing import TypedDict, Optional


class CharacterCreation(BaseModel):
    """Schema for creating a new character"""
    name: str = Field(..., description="The character's name")
    gender: str = Field(..., description="The character's gender")
    race: str = Field(..., description="The character's race")
    alignment: Optional[str] = Field(None, description="The character's moral and ethical alignment")

    appearance: Optional[str] = Field(None, description="The character's appearance description")
    personality: Optional[str] = Field(None, description="The character's personality traits")
    backstory: Optional[str] = Field(None, description="The character's backstory")

    universe: str = Field(..., description="The universe the character belongs to")
    world_theme: str = Field(..., description="The theme of the world")
    tone: str = Field(..., description="The tone of the character's story")

    linked_world_id: Optional[UUID] = Field(None, description="ID of a linked world, if any")

    class Config:
        json_schema_extra = {
            "example": {
                "name": "Elara Nightshade",
                "gender": "Female",
                "race": "Half-Elf",
                "alignment": "Chaotic Good",
                "appearance": "Tall with silver hair and violet eyes",
                "personality": "Mysterious and calculating",
                "backstory": "Born in the shadow of the Misty Mountains...",
                "universe": "Fantasy",
                "world_theme": "High Fantasy",
                "tone": "Serious",
                "linked_world_id": None
            },
            "description": "Schema for creating a new character with all attributes"
        }


class Character(BaseModel):
    """Schema for a character with all attributes"""
    id: UUID = Field(..., description="Unique identifier for the character")
    name: str = Field(..., description="The character's name")
    gender: str = Field(..., description="The character's gender")
    race: str = Field(..., description="The character's race")
    alignment: str = Field(..., description="The character's moral and ethical alignment")

    appearance: str = Field(..., description="The character's appearance description")
    personality: str = Field(..., description="The character's personality traits")
    backstory: str = Field(..., description="The character's backstory")

    universe: str = Field(..., description="The universe the character belongs to")
    world_theme: str = Field(..., description="The theme of the world")
    tone: str = Field(..., description="The tone of the character's story")

    linked_world_id: Optional[UUID] = Field(None, description="ID of a linked world, if any")
    image_filename: Optional[str] = Field(None, description="Filename of the character's image, if any")

    class Config:
        json_schema_extra = {
            "example": {
                "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
                "name": "Elara Nightshade",
                "gender": "Female",
                "race": "Half-Elf",
                "alignment": "Chaotic Good",
                "appearance": "Tall with silver hair and violet eyes",
                "personality": "Mysterious and calculating",
                "backstory": "Born in the shadow of the Misty Mountains...",
                "universe": "Fantasy",
                "world_theme": "High Fantasy",
                "tone": "Serious",
                "linked_world_id": None,
                "image_filename": "character_image_3fa85f64-5717-4562-b3fc-2c963f66afa6.png"
            }
        }
