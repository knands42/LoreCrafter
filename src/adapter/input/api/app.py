from dotenv import load_dotenv

load_dotenv()

from pathlib import Path

from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from fastapi.staticfiles import StaticFiles

from src.adapter.input.api.routers import character_router, world_router, campaign_router, asset_router, pdf_router, users_router

APP_TITLE = "LoreCrafter API"
APP_DESCRIPTION = "API for LoreCrafter, a tool that helps users create rich backstories for tabletop role-playing game characters, worlds, and campaigns using AI."
APP_VERSION = "1.0.0"

# Create FastAPI app
app = FastAPI(
    title=APP_TITLE,
    description=APP_DESCRIPTION,
    version=APP_VERSION,
    docs_url="/docs",
    redoc_url="/redoc",
    openapi_tags=[
        {
            "name": "characters",
            "description": "Operations with characters. Create, retrieve, and search for characters with AI-generated backstories."
        },
        {
            "name": "worlds",
            "description": "Operations with worlds. Create, retrieve, and search for worlds with AI-generated histories and timelines."
        },
        {
            "name": "campaigns",
            "description": "Operations with campaigns. Create and retrieve campaigns with AI-generated settings and hidden elements."
        },
        {
            "name": "assets",
            "description": "Operations with assets. Retrieve character and world images."
        },
        {
            "name": "pdf",
            "description": "Operations with PDFs. Generate PDF documents for characters and worlds."
        },
        {
            "name": "users",
            "description": "Operations with users. Create users, sign in, and manage authentication."
        },
    ]
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
app.include_router(users_router.router, prefix="/api", tags=["users"])

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
async def healthcheck():
    return {"message": "API is healthy"}
