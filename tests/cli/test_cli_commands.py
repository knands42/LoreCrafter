"""
Tests for CLI commands.
"""
import json
import uuid
from unittest.mock import patch, MagicMock

import pytest
from typer.testing import CliRunner

from main import app
from src.application.domain.character_domain import CharacterDomain
from src.application.domain.word_domain import WorldDomain
from src.application.domain.campaign_domain import CampaignDomain


@pytest.fixture
def cli_runner():
    """
    Returns a CliRunner instance for testing Typer CLI commands.
    """
    return CliRunner()


@pytest.fixture
def mock_character_generator():
    """
    Mock for CharacterGenerator.
    """
    with patch("main.CharacterGenerator") as mock:
        generator_instance = mock.return_value
        # Mock the generate method to return a predefined character
        generator_instance.generate.return_value = {
            "id": str(uuid.uuid4()),
            "name": "Test Character",
            "gender": "Male",
            "race": "Human",
            "appearance": "Tall with brown hair",
            "personality": "Brave and kind",
            "backstory": "Born in a small village",
            "universe": "Fantasy",
            "world_theme": "Medieval",
            "tone": "Heroic",
            "linked_world_id": None,
            "image_filename": None
        }
        yield generator_instance


@pytest.fixture
def mock_world_generator():
    """
    Mock for WorldGenerator.
    """
    with patch("main.WorldGenerator") as mock:
        generator_instance = mock.return_value
        # Mock the generate method to return a predefined world
        generator_instance.generate.return_value = {
            "id": str(uuid.uuid4()),
            "name": "Test World",
            "universe": "Fantasy",
            "world_theme": "Medieval",
            "tone": "Epic",
            "backstory": "A world of magic and adventure",
            "timeline": "Ancient era -> Medieval era -> Modern era",
            "image_filename": None
        }
        yield generator_instance


@pytest.fixture
def mock_campaign_generator():
    """
    Mock for CampaignGenerator.
    """
    with patch("main.CampaignGenerator") as mock:
        generator_instance = mock.return_value
        # Mock the generate method to return a predefined campaign
        generator_instance.generate.return_value = {
            "id": str(uuid.uuid4()),
            "name": "Test Campaign",
            "universe": "Fantasy",
            "world_theme": "Medieval",
            "tone": "Epic",
            "campaign": "A grand adventure to save the kingdom",
            "hidden_elements": "Secret dragon cult",
            "linked_world_id": None
        }
        yield generator_instance


@pytest.fixture
def mock_vector_store():
    """
    Mock for VectorStore.
    """
    with patch("main.CharacterVectorStore") as char_mock, \
         patch("main.WorldVectorStore") as world_mock:
        
        # Mock character search
        char_store = char_mock.return_value
        mock_char_doc = MagicMock()
        mock_char_doc.page_content = json.dumps({
            "id": str(uuid.uuid4()),
            "name": "Found Character",
            "gender": "Female",
            "race": "Elf",
            "appearance": "Slender with blonde hair",
            "personality": "Wise and mysterious",
            "backstory": "Lived for centuries in an enchanted forest",
            "universe": "Fantasy",
            "world_theme": "Medieval",
            "tone": "Mystical",
            "linked_world_id": None,
            "image_filename": None
        })
        char_store.search_similar.return_value = [mock_char_doc]
        
        # Mock world search
        world_store = world_mock.return_value
        mock_world_doc = MagicMock()
        mock_world_doc.page_content = json.dumps({
            "id": str(uuid.uuid4()),
            "name": "Found World",
            "universe": "Fantasy",
            "world_theme": "Medieval",
            "tone": "Epic",
            "backstory": "A world of magic and adventure",
            "timeline": "Ancient era -> Medieval era -> Modern era",
            "image_filename": None
        })
        world_store.search_similar.return_value = [mock_world_doc]
        
        yield char_store, world_store


def test_create_character_default(cli_runner, mock_character_generator):
    """
    Test creating a character with default values.
    """
    with patch("main.create_character_pdf") as mock_pdf:
        result = cli_runner.invoke(app, ["create-character", "--default"])
        
        assert result.exit_code == 0
        assert "Welcome to LoreCrafter" in result.stdout
        assert mock_character_generator.generate.called
        assert mock_pdf.called


def test_search_character(cli_runner, mock_vector_store):
    """
    Test searching for a character.
    """
    char_store, _ = mock_vector_store
    
    result = cli_runner.invoke(app, ["search-character", "elf"])
    
    assert result.exit_code == 0
    assert char_store.search_similar.called
    assert "Found Character" in result.stdout


def test_create_world_default(cli_runner, mock_world_generator):
    """
    Test creating a world with default values.
    """
    with patch("main.create_world_pdf") as mock_pdf:
        result = cli_runner.invoke(app, ["create-world", "--default"])
        
        assert result.exit_code == 0
        assert "Welcome to LoreCrafter World Builder" in result.stdout
        assert mock_world_generator.generate.called
        assert mock_pdf.called


def test_search_world(cli_runner, mock_vector_store):
    """
    Test searching for a world.
    """
    _, world_store = mock_vector_store
    
    result = cli_runner.invoke(app, ["search-world", "fantasy"])
    
    assert result.exit_code == 0
    assert world_store.search_similar.called
    assert "Found World" in result.stdout


def test_create_campaign_default(cli_runner, mock_campaign_generator):
    """
    Test creating a campaign with default values.
    """
    result = cli_runner.invoke(app, ["create-campaign", "--default"])
    
    assert result.exit_code == 0
    assert "Welcome to LoreCrafter Campaign Builder" in result.stdout
    assert mock_campaign_generator.generate.called


def test_api_command(cli_runner):
    """
    Test the API command (just check it starts without errors).
    """
    with patch("main.uvicorn.run") as mock_run:
        result = cli_runner.invoke(app, ["api", "--host", "127.0.0.1", "--port", "8000"])
        
        assert result.exit_code == 0
        assert "Starting LoreCrafter API server" in result.stdout
        assert mock_run.called
        mock_run.assert_called_with(
            "src.adapter.input.api.app:app",
            host="127.0.0.1",
            port=8000,
            reload=False
        )