from fastapi import APIRouter, HTTPException, Depends
from typing import List, Optional
import json
from uuid import UUID

from src.adapter.output.repository import CharacterVectorStore
from src.application.domain.character_domain import CharacterCreation, Character
from src.application.usecases import CharacterGenerator

# Create router
router = APIRouter()

# Dependencies
def get_character_vector_store():
    return CharacterVectorStore()

def get_character_generator(vector_store: CharacterVectorStore = Depends(get_character_vector_store)):
    return CharacterGenerator(vector_store)

# API endpoints
@router.post("/characters", response_model=Character)
async def create_character(
    character_info: CharacterCreation,
    character_generator: CharacterGenerator = Depends(get_character_generator)
):
    """
    Create a new character with AI-generated backstory, personality, and appearance.

    This endpoint takes basic character information and uses AI to generate a rich backstory,
    personality traits, and detailed appearance description.
    """
    try:
        return character_generator.generate(character_info)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error creating character: {str(e)}")

@router.get("/characters/search", response_model=List[Character])
async def search_characters(
    query: str,
    top: Optional[int] = 2,
    vector_store: CharacterVectorStore = Depends(get_character_vector_store)
):
    """
    Search for characters based on a query string.

    This endpoint performs a semantic search on character data and returns the most relevant matches.
    The 'top' parameter controls how many results to return.
    """
    # TODO: Implement search functionality
    try:
        return []
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error searching characters: {str(e)}")

@router.get("/characters/{character_id}", response_model=Character)
async def get_character(
    character_id: UUID,
    vector_store: CharacterVectorStore = Depends(get_character_vector_store)
):
    """
    Retrieve a specific character by ID.

    This endpoint returns detailed information about a character, including its backstory,
    personality, appearance, and other attributes.
    """
    # TODO: Implement retrieval functionality
    try:
        return None
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving character: {str(e)}")
