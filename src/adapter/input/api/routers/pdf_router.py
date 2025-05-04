from fastapi import APIRouter, HTTPException, Depends
from fastapi.responses import FileResponse
from pathlib import Path
import json
from uuid import UUID
import tempfile

from src.adapter.output.repository import CharacterVectorStore, WorldVectorStore
from src.adapter.output.pdf import create_character_pdf, create_world_pdf

# Create router
router = APIRouter()

# Dependencies
def get_character_vector_store():
    return CharacterVectorStore()

def get_world_vector_store():
    return WorldVectorStore()

# API endpoints
@router.get("/pdf/character/{character_id}")
async def generate_character_pdf(
    character_id: UUID,
    vector_store: CharacterVectorStore = Depends(get_character_vector_store)
):
    """
    Generate a PDF for a character.
    """
    try:
        # Search for the character by ID
        results = vector_store.search_similar(f"id:{character_id}", 1)
        if not results:
            raise HTTPException(status_code=404, detail=f"Character with ID {character_id} not found")
        
        character = json.loads(results[0].page_content)
        
        # Create a temporary file for the PDF
        with tempfile.NamedTemporaryFile(suffix=".pdf", delete=False) as tmp:
            pdf_path = tmp.name
        
        # Generate the PDF
        create_character_pdf(character, pdf_path)
        
        # Return the PDF file
        return FileResponse(
            path=pdf_path,
            filename=f"character_{character_id}.pdf",
            media_type="application/pdf"
        )
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error generating character PDF: {str(e)}")

@router.get("/pdf/world/{world_id}")
async def generate_world_pdf(
    world_id: UUID,
    vector_store: WorldVectorStore = Depends(get_world_vector_store)
):
    """
    Generate a PDF for a world.
    """
    try:
        # Search for the world by ID
        results = vector_store.search_similar(f"id:{world_id}", 1)
        if not results:
            raise HTTPException(status_code=404, detail=f"World with ID {world_id} not found")
        
        world = json.loads(results[0].page_content)
        
        # Create a temporary file for the PDF
        with tempfile.NamedTemporaryFile(suffix=".pdf", delete=False) as tmp:
            pdf_path = tmp.name
        
        # Generate the PDF
        create_world_pdf(world, pdf_path)
        
        # Return the PDF file
        return FileResponse(
            path=pdf_path,
            filename=f"world_{world_id}.pdf",
            media_type="application/pdf"
        )
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error generating world PDF: {str(e)}")