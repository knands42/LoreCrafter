from pathlib import Path

from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles

from src.adapter.input.api.routers import character_router, world_router, campaign_router, asset_router, pdf_router

# Create FastAPI app
app = FastAPI(
    title="LoreCrafter API",
    description="API for LoreCrafter, a tool that helps users create rich backstories for tabletop role-playing game characters, worlds, and campaigns using AI.",
    version="1.0.0",
)

# Add CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include routers
app.include_router(character_router.router, prefix="/api", tags=["characters"])
app.include_router(world_router.router, prefix="/api", tags=["worlds"])
app.include_router(campaign_router.router, prefix="/api", tags=["campaigns"])
app.include_router(asset_router.router, prefix="/api", tags=["assets"])
app.include_router(pdf_router.router, prefix="/api", tags=["pdf"])

# Mount assets directory
assets_dir = Path("assets")
assets_dir.mkdir(exist_ok=True)
app.mount("/assets", StaticFiles(directory=assets_dir), name="assets")


# Error handling
@app.exception_handler(HTTPException)
async def http_exception_handler(request, exc):
    return {
        "error": exc.detail,
        "status_code": exc.status_code
    }


@app.get("/")
async def root():
    return {"message": "Welcome to LoreCrafter API"}


@app.get("/healthcheck")
async def root():
    return {"message": "API is healthy"}
