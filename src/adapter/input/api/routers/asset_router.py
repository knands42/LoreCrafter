from fastapi import APIRouter, HTTPException
from fastapi.responses import FileResponse
from pathlib import Path

# Create router
router = APIRouter()

# API endpoints
@router.get("/assets/character/{image_filename}")
async def get_character_image(image_filename: str):
    """
    Retrieve a character image by filename.
    """
    try:
        image_path = Path("assets") / image_filename
        if not image_path.exists():
            raise HTTPException(status_code=404, detail=f"Image {image_filename} not found")
        
        return FileResponse(image_path)
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving character image: {str(e)}")

@router.get("/assets/world/{image_filename}")
async def get_world_image(image_filename: str):
    """
    Retrieve a world image by filename.
    """
    try:
        image_path = Path("assets") / image_filename
        if not image_path.exists():
            raise HTTPException(status_code=404, detail=f"Image {image_filename} not found")
        
        return FileResponse(image_path)
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving world image: {str(e)}")