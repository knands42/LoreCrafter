from fastapi import APIRouter, HTTPException, Depends
from typing import List, Optional
import json
from uuid import UUID

from src.adapter.output.repository import WorldVectorStore
from src.application.usecases import CampaignGenerator
from src.application.domain.campaign_domain import CampaignCreationDomain, CampaignDomain

# Create router
router = APIRouter()

# Dependencies
def get_world_vector_store():
    return WorldVectorStore()

def get_campaign_generator(vector_store: WorldVectorStore = Depends(get_world_vector_store)):
    return CampaignGenerator(vector_store)

# API endpoints
@router.post("/campaigns", response_model=CampaignDomain)
async def create_campaign(
    campaign_info: CampaignCreationDomain,
    campaign_generator: CampaignGenerator = Depends(get_campaign_generator)
):
    """
    Create a new campaign with AI-generated setting and hidden elements.
    """
    try:
        result = campaign_generator.generate(campaign_info)
        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error creating campaign: {str(e)}")

@router.get("/campaigns/{campaign_id}", response_model=CampaignDomain)
async def get_campaign(
    campaign_id: UUID,
    vector_store: WorldVectorStore = Depends(get_world_vector_store)
):
    """
    Retrieve a specific campaign by ID.
    """
    try:
        # Search for the campaign by ID
        # Note: This assumes campaigns are stored in the world vector store
        # You may need to adjust this based on your actual implementation
        results = vector_store.search_similar(f"id:{campaign_id}", 1)
        if not results:
            raise HTTPException(status_code=404, detail=f"Campaign with ID {campaign_id} not found")
        
        campaign = json.loads(results[0].page_content)
        return campaign
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error retrieving campaign: {str(e)}")