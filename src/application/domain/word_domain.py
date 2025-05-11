from pydantic import BaseModel, Field
from typing import Optional
from uuid import UUID


class WorldCreation(BaseModel):
    """Schema for creating a new world"""
    name: str = Field(..., description="The world's name")
    universe: str = Field(..., description="The universe the world belongs to")
    world_theme: str = Field(..., description="The theme of the world")
    tone: str = Field(..., description="The tone of the world's story")
    backstory: Optional[str] = Field(None, description="The world's backstory")
    timeline: Optional[str] = Field(None, description="The world's timeline of events")

    class Config:
        json_schema_extra = {
            "example": {
                "name": "Eldoria",
                "universe": "Fantasy",
                "world_theme": "High Fantasy",
                "tone": "Epic",
                "backstory": "A realm of ancient magic and forgotten kingdoms...",
                "timeline": "Age of Creation: The gods forged Eldoria from the void..."
            }
        }


class World(BaseModel):
    """Schema for a world with all attributes"""
    id: UUID = Field(..., description="Unique identifier for the world")
    name: str = Field(..., description="The world's name")
    universe: str = Field(..., description="The universe the world belongs to")
    world_theme: str = Field(..., description="The theme of the world")
    tone: str = Field(..., description="The tone of the world's story")
    backstory: str = Field(..., description="The world's backstory")
    timeline: str = Field(..., description="The world's timeline of events")
    image_filename: Optional[str] = Field(None, description="Filename of the world's image, if any")

    class Config:
        json_schema_extra = {
            "example": {
                "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
                "name": "Eldoria",
                "universe": "Fantasy",
                "world_theme": "High Fantasy",
                "tone": "Epic",
                "backstory": "A realm of ancient magic and forgotten kingdoms...",
                "timeline": "Age of Creation: The gods forged Eldoria from the void...",
                "image_filename": "world_image_3fa85f64-5717-4562-b3fc-2c963f66afa6.png"
            }
        }
