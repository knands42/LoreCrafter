"""
Tests for basic API endpoints.
"""
import pytest
from fastapi.testclient import TestClient


def test_root_endpoint(api_client):
    """
    Test that the root endpoint returns the expected response.
    """
    response = api_client.get("/")
    assert response.status_code == 200
    assert response.json() == {"message": "Welcome to LoreCrafter API"}


def test_healthcheck_endpoint(api_client):
    """
    Test that the healthcheck endpoint returns the expected response.
    """
    response = api_client.get("/healthcheck")
    assert response.status_code == 200
    assert response.json() == {"message": "API is healthy"}