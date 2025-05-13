"""
Integration tests for user API endpoints.
"""
from datetime import datetime, UTC

from src.adapter.output.sql_db.models.user_model import UserModel


def test_create_user(api_client):
    """
    Test creating a new user.
    """
    # Given
    user_data = {
        "username": "testuser",
        "email": "test@example.com",
        "password": "securepassword"
    }

    # When
    response = api_client.post("/api/users", json=user_data)

    # Then
    assert response.status_code == 200
    result = response.json()
    assert result["username"] == user_data["username"]
    assert result["email"] == user_data["email"]
    assert "password" not in result
    assert result["is_active"] is True

    user = UserModel.get(UserModel.username == user_data["username"])
    assert user.username == user_data["username"]
    assert user.email == user_data["email"]
    assert user.password_hash is not None
    assert user.is_active is True


def test_create_user_existing_username(api_client):
    """
    Test creating a user with an existing username.
    """
    # Create an existing user
    existing_user = UserModel.create(
        username="existinguser",
        email="existing@example.com",
        password_hash="hashed_password",
        is_active=True,
        created_at=datetime.now(UTC),
        updated_at=datetime.now(UTC)
    )

    # Create user data with existing username
    user_data = {
        "username": "existinguser",
        "email": "new@example.com",
        "password": "securepassword"
    }

    # Send request to create user
    response = api_client.post("/api/users", json=user_data)

    # Check response
    assert response.status_code == 400
    assert "Username already exists" in response.json()["detail"]

    # Verify no new user was created
    users = list(UserModel.select().where(UserModel.username == "existinguser"))
    assert len(users) == 1
    assert users[0].email == "existing@example.com"  # Original email, not the new one


def test_signin_success(api_client):
    """
    Test successful user signin.
    """
    # Create a test user with a known password
    password = "correctpassword"

    # Create a user through the API to ensure proper password hashing
    user_data = {
        "username": "testuser",
        "email": "test@example.com",
        "password": password
    }

    # Create the user
    api_client.post("/api/users", json=user_data)

    # Send signin request
    response = api_client.post(
        "/api/signin",
        data={"username": "testuser", "password": password},
        headers={"Content-Type": "application/x-www-form-urlencoded"}
    )

    # Check response
    assert response.status_code == 200
    result = response.json()
    assert "access_token" in result
    assert result["token_type"] == "bearer"

    # Check that a cookie was set
    cookies = response.cookies
    assert "access_token" in cookies


def test_signin_invalid_credentials(api_client):
    """
    Test signin with invalid credentials.
    """
    # Create a test user with a known password
    correct_password = "correctpassword"
    wrong_password = "wrongpassword"

    # Create a user through the API to ensure proper password hashing
    user_data = {
        "username": "testuser2",
        "email": "test2@example.com",
        "password": correct_password
    }

    # Create the user
    api_client.post("/api/users", json=user_data)

    # Send signin request with wrong password
    response = api_client.post(
        "/api/signin",
        data={"username": "testuser2", "password": wrong_password},
        headers={"Content-Type": "application/x-www-form-urlencoded"}
    )

    # Check response
    assert response.status_code == 401
    assert "Invalid username or password" in response.json()["detail"]


def test_signin_user_not_found(api_client):
    """
    Test signin with a non-existent user.
    """
    # Send signin request with non-existent username
    response = api_client.post(
        "/api/signin",
        data={"username": "nonexistentuser", "password": "anypassword"},
        headers={"Content-Type": "application/x-www-form-urlencoded"}
    )

    # Check response
    assert response.status_code == 401
    assert "Invalid username or password" in response.json()["detail"]


def test_get_current_user(api_client):
    """
    Test getting the current user with a valid token.
    """
    # Create a test user
    user_data = {
        "username": "testuser3",
        "email": "test3@example.com",
        "password": "securepassword"
    }

    # Create the user
    api_client.post("/api/users", json=user_data)

    # Sign in to get a token
    signin_response = api_client.post(
        "/api/signin",
        data={"username": "testuser3", "password": "securepassword"},
        headers={"Content-Type": "application/x-www-form-urlencoded"}
    )

    # Get the token from the response
    token = signin_response.json()["access_token"]

    # Send request with token cookie
    response = api_client.get(
        "/api/me",
        cookies={"access_token": token}
    )

    # Check response
    assert response.status_code == 200
    result = response.json()
    assert result["username"] == "testuser3"
    assert result["email"] == "test3@example.com"
    assert result["is_active"] is True


def test_get_current_user_no_token(api_client, test_db):
    """
    Test getting the current user without a token.
    """
    # Send request without token cookie
    response = api_client.get("/api/me")

    # Check response
    assert response.status_code == 401
    assert "Not authenticated" in response.json()["detail"]


def test_get_current_user_invalid_token(api_client):
    """
    Test getting the current user with an invalid token.
    """
    # Send request with invalid token cookie
    response = api_client.get(
        "/api/me",
        cookies={"access_token": "invalid_token"}
    )

    # Check response
    assert response.status_code == 401
    assert "Invalid token" in response.json()["detail"]


def test_get_current_user_not_found(api_client):
    """
    Test getting a user that doesn't exist.
    """
    # Create a test user
    user_data = {
        "username": "tempuser",
        "email": "temp@example.com",
        "password": "securepassword"
    }

    # Create the user
    api_client.post("/api/users", json=user_data)

    # Sign in to get a token
    signin_response = api_client.post(
        "/api/signin",
        data={"username": "tempuser", "password": "securepassword"},
        headers={"Content-Type": "application/x-www-form-urlencoded"}
    )

    # Get the token from the response
    token = signin_response.json()["access_token"]

    # Delete the user from the database
    user = UserModel.get(UserModel.username == "tempuser")
    user.delete_instance()

    # Send request with token cookie
    response = api_client.get(
        "/api/me",
        cookies={"access_token": token}
    )

    # Check response
    assert response.status_code == 404
    assert "User not found" in response.json()["detail"]
