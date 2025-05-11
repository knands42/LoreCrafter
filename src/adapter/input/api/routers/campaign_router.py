import json
from uuid import UUID

from fastapi import APIRouter, HTTPException, Depends

from src.adapter.output.repository import WorldVectorStore
from src.application.domain.campaign_domain import CampaignCreation, Campaign
from src.application.usecases import CampaignGenerator

# Create router
router = APIRouter()

# Dependencies
def get_world_vector_store():
    return WorldVectorStore()

def get_campaign_generator(vector_store: WorldVectorStore = Depends(get_world_vector_store)):
    return CampaignGenerator(vector_store)

# API endpoints
@router.post("/campaigns", response_model=Campaign)
async def create_campaign(
    campaign_info: CampaignCreation,
    campaign_generator: CampaignGenerator = Depends(get_campaign_generator)
):
    """
    Create a new campaign with AI-generated setting and hidden elements.

    This endpoint takes basic campaign information and uses AI to generate a rich campaign setting
    and hidden elements that can be used by the game master.
    """
    try:
        return campaign_generator.generate(campaign_info)
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error creating campaign: {str(e)}")

@router.get("/campaigns/{campaign_id}", response_model=Campaign)
async def get_campaign(
    campaign_id: UUID,
    vector_store: WorldVectorStore = Depends(get_world_vector_store)
):
    """
    Retrieve a specific campaign by ID.

    This endpoint returns detailed information about a campaign, including its setting,
    hidden elements, and other attributes.
    """
    # TODO: Implement retrieval functionality
    try:
        return None
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving campaign: {str(e)}")
