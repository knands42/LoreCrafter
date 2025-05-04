import uuid
from typing import TypedDict


class CharacterDomain(TypedDict):
    id: uuid.UUID
    name: str
    gender: str
    race: str

    appearance: str
    personality: str
    backstory: str

    universe: str
    world_theme: str
    tone: str

    linked_world_id: uuid.UUID | None
    image_filename: str | None


class CharacterCreateDomain(TypedDict):
    name: str
    gender: str
    race: str

    appearance: str | None
    personality: str | None
    backstory: str | None

    universe: str
    world_theme: str
    tone: str

    linked_world_id: uuid.UUID | None
