"""
Tests for character API endpoints.
"""
import uuid
from unittest.mock import patch

import pytest

from src.adapter.output.vector_db.character_vector_store import CharacterVectorStore
from src.application.usecases import CharacterGenerator


@pytest.fixture
def spy_character_generator():
    with patch("src.application.usecases.CharacterGenerator", wraps=CharacterGenerator) as spy:
        yield spy


@pytest.fixture
def spy_vector_store():
    with patch("src.adapter.input.api.routers.character_router.CharacterVectorStore", wrap=CharacterVectorStore) as spy:
        yield spy


def test_create_character_full(api_client, character_name="Elina"):
    character_data = {
        "name": character_name,
        "gender": "Female",
        "race": "Dwarf",
        "appearance": "Pale face with emerald green eyes",
        "personality": "Stubborn and unpredictable",
        "backstory": "A fallen prince trying restore his legacy",
        "universe": "Fantasy",
        "world_theme": "Medieval",
        "tone": "Gritty",
        "linked_world_id": None
    }

    response = api_client.post("/api/characters", json=character_data)

    assert response.status_code == 200

    result = response.json()
    assert result["name"] == character_data.get("name")
    assert result["gender"] == character_data.get("gender")
    assert result["race"] == character_data.get("race")

    assert result["universe"] == character_data.get("universe")
    assert result["world_theme"] == character_data.get("world_theme")
    assert result["tone"] == character_data.get("tone")
    assert result["linked_world_id"] == character_data.get("linked_world_id")
    assert result['image_filename'] is not None

    assert "appearance" in result
    assert "personality" in result
    assert "backstory" in result

    assert result["appearance"] != character_data.get("appearance")
    assert result["personality"] != character_data.get("personality")
    assert result["backstory"] != character_data.get("backstory")


def test_create_character_partial(api_client):
    character_data = {
        "name": "Orkrid",
        "gender": "Male",
        "race": "Orc",
        "appearance": None,
        "personality": None,
        "backstory": None,
        "universe": "Fantasy",
        "world_theme": "Medieval",
        "tone": "Gritty",
        "linked_world_id": None
    }

    response = api_client.post("/api/characters", json=character_data)

    assert response.status_code == 200

    result = response.json()
    assert result["name"] == character_data.get("name")
    assert result["gender"] == character_data.get("gender")
    assert result["race"] == character_data.get("race")

    assert result["universe"] == character_data.get("universe")
    assert result["world_theme"] == character_data.get("world_theme")
    assert result["tone"] == character_data.get("tone")
    assert result["linked_world_id"] == character_data.get("linked_world_id")
    assert result['image_filename'] is not None

    assert "appearance" in result
    assert "personality" in result
    assert "backstory" in result


def test_search_characters(api_client):
    """
    Test searching for characters.
    """
    test_create_character_full(api_client, character_name="Melina")
    response = api_client.get("/api/characters/search", params={"query": "A female dwarf character called Melina", "top": 1})

    assert response.status_code == 200

    results = response.json()
    assert len(results) == 1
    assert results[0]["name"] == "Melina"
    assert results[0]["gender"] == "Female"
    assert results[0]["race"] == "Dwarf"


def test_get_character_by_id(api_client):
    """
    Test getting a character by ID.
    """
    search_response = api_client.get("/api/characters/search", params={"query": "A female dwarf character called Elina", "top": 1})
    character_id = search_response.json()[0]["id"]
    
    response = api_client.get(f"/api/characters/{character_id}")
    assert response.status_code == 200

    result = response.json()
    assert result["name"] == "Elina"
    assert result["gender"] == "Female"
    assert result["race"] == "Dwarf"

def test_get_character_not_found(api_client):
    """
    Test getting a character that doesn't exist.
    """
    character_id = uuid.uuid4()
    response = api_client.get(f"/api/characters/{character_id}")

    assert response.status_code == 404
    assert "not found" in response.json()["detail"]
