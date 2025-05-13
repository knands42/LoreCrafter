from fastapi import APIRouter, HTTPException, Depends
from typing import List, Optional
from uuid import UUID

from src.adapter.output.vector_db.world_vector_store import WorldVectorStore
from src.application.domain.word_domain import WorldCreation, World
from src.application.usecases import WorldGenerator

# Create router
router = APIRouter()

# Dependencies
def get_world_vector_store():
    return WorldVectorStore()

def get_world_generator(vector_store: WorldVectorStore = Depends(get_world_vector_store)):
    return WorldGenerator(vector_store)

# API endpoints
@router.post("/worlds", response_model=World)
async def create_world(
    world_info: WorldCreation,
    world_generator: WorldGenerator = Depends(get_world_generator)
):
    """
    Create a new world with AI-generated history and timeline.

    This endpoint takes basic world information and uses AI to generate a rich backstory
    and detailed timeline of events for the world.
    """
    try:
        return world_generator.generate(world_info)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error creating world: {str(e)}")

@router.get("/worlds/search", response_model=List[World])
async def search_worlds(
    query: str,
    top: Optional[int] = 2,
    vector_store: WorldVectorStore = Depends(get_world_vector_store)
):
    """
    Search for worlds based on a query string.

    This endpoint performs a semantic search on world data and returns the most relevant matches.
    The 'top' parameter controls how many results to return.
    """
    # TODO: Implement search functionality
    try:
        return []
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error searching worlds: {str(e)}")

@router.get("/worlds/{world_id}", response_model=World)
async def get_world(
    world_id: UUID,
    vector_store: WorldVectorStore = Depends(get_world_vector_store)
):
    """
    Retrieve a specific world by ID.

    This endpoint returns detailed information about a world, including its backstory,
    timeline, and other attributes.
    """
    # TODO: Implement retrieval functionality
    try:
        return None
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving world: {str(e)}")
