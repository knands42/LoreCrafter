from fastapi import APIRouter, HTTPException, Depends
from typing import List, Optional
import json
from uuid import UUID

from src.adapter.output.repository import WorldVectorStore
from src.application.usecases import WorldGenerator
from src.application.domain.word_domain import WorldCreateDomain, WorldDomain

# Create router
router = APIRouter()

# Dependencies
def get_world_vector_store():
    return WorldVectorStore()

def get_world_generator(vector_store: WorldVectorStore = Depends(get_world_vector_store)):
    return WorldGenerator(vector_store)

# API endpoints
@router.post("/worlds", response_model=WorldDomain)
async def create_world(
    world_info: WorldCreateDomain,
    world_generator: WorldGenerator = Depends(get_world_generator)
):
    """
    Create a new world with AI-generated history and timeline.
    """
    try:
        result = world_generator.generate(world_info)
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error creating world: {str(e)}")

@router.get("/worlds/search", response_model=List[WorldDomain])
async def search_worlds(
    query: str,
    top: Optional[int] = 2,
    vector_store: WorldVectorStore = Depends(get_world_vector_store)
):
    """
    Search for worlds based on a query string.
    """
    try:
        results = vector_store.search_similar(query, top)
        worlds = []
        for doc in results:
            world = json.loads(doc.page_content)
            worlds.append(world)
        return worlds
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error searching worlds: {str(e)}")

@router.get("/worlds/{world_id}", response_model=WorldDomain)
async def get_world(
    world_id: UUID,
    vector_store: WorldVectorStore = Depends(get_world_vector_store)
):
    """
    Retrieve a specific world by ID.
    """
    try:
        # Search for the world by ID
        results = vector_store.search_similar(f"id:{world_id}", 1)
        if not results:
            raise HTTPException(status_code=404, detail=f"World with ID {world_id} not found")
        
        world = json.loads(results[0].page_content)
        return world
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving world: {str(e)}")