from fastapi import APIRouter, HTTPException, Depends
from typing import List, Optional
import json
from uuid import UUID

from src.adapter.output.repository import CharacterVectorStore
from src.application.usecases import CharacterGenerator
from src.application.domain.character_domain import CharacterCreateDomain, CharacterDomain

# Create router
router = APIRouter()

# Dependencies
def get_character_vector_store():
    return CharacterVectorStore()

def get_character_generator(vector_store: CharacterVectorStore = Depends(get_character_vector_store)):
    return CharacterGenerator(vector_store)

# API endpoints
@router.post("/characters", response_model=CharacterDomain)
async def create_character(
    character_info: CharacterCreateDomain,
    character_generator: CharacterGenerator = Depends(get_character_generator)
):
    """
    Create a new character with AI-generated backstory, personality, and appearance.
    """
    try:
        result = character_generator.generate(character_info)
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error creating character: {str(e)}")

@router.get("/characters/search", response_model=List[CharacterDomain])
async def search_characters(
    query: str,
    top: Optional[int] = 2,
    vector_store: CharacterVectorStore = Depends(get_character_vector_store)
):
    """
    Search for characters based on a query string.
    """
    try:
        results = vector_store.search_similar(query, top)
        characters = []
        for doc in results:
            character = json.loads(doc.page_content)
            characters.append(character)
        return characters
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error searching characters: {str(e)}")

@router.get("/characters/{character_id}", response_model=CharacterDomain)
async def get_character(
    character_id: UUID,
    vector_store: CharacterVectorStore = Depends(get_character_vector_store)
):
    """
    Retrieve a specific character by ID.
    """
    pass
