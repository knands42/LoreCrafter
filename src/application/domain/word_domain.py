import uuid
from typing import TypedDict


class WorldDomain(TypedDict):
    id: uuid.UUID
    name: str
    universe: str
    world_theme: str
    tone: str
    backstory: str
    timeline: str
    image_filename: str | None


class WorldCreateDomain(TypedDict):
    name: str
    universe: str
    world_theme: str
    tone: str
    backstory: str | None
    timeline: str | None
