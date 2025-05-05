"""
This file contains pytest fixtures that can be used across all tests.
"""
import os
import pytest
from fastapi.testclient import TestClient

# Import the FastAPI app
from src.adapter.input.api.app import app


@pytest.fixture
def api_client():
    return TestClient(app)


@pytest.fixture
def mock_env_vars(monkeypatch):
    # monkeypatch.setenv("OPENAI_API_KEY", "test_api_key")
    # monkeypatch.setenv("GOOGLE_API_KEY", "test_api_key")
    pass
