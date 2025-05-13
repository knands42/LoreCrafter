from datetime import datetime
from typing import Optional
from uuid import UUID

from pydantic import BaseModel, Field, EmailStr


class UserCreation(BaseModel):
    """Schema for creating a new user"""
    username: str = Field(..., description="The user's username", min_length=3, max_length=80)
    email: EmailStr = Field(..., description="The user's email address", max_length=120)
    password: str = Field(..., description="The user's password", min_length=8)

    class Config:
        json_schema_extra = {
            "example": {
                "username": "johndoe",
                "email": "john.doe@example.com",
                "password": "securepassword123"
            },
            "description": "Schema for creating a new user"
        }


class User(BaseModel):
    """Schema for a user with all attributes"""
    id: UUID = Field(..., description="Unique identifier for the user")
    username: str = Field(..., description="The user's username")
    email: EmailStr = Field(..., description="The user's email address")
    is_active: bool = Field(..., description="Whether the user is active")
    created_at: datetime = Field(..., description="When the user was created")
    updated_at: datetime = Field(..., description="When the user was last updated")

    class Config:
        json_schema_extra = {
            "example": {
                "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
                "username": "johndoe",
                "email": "john.doe@example.com",
                "is_active": True,
                "created_at": "2023-01-01T00:00:00Z",
                "updated_at": "2023-01-01T00:00:00Z"
            }
        }


class UserLogin(BaseModel):
    """Schema for user login"""
    username: str = Field(..., description="The user's username")
    password: str = Field(..., description="The user's password")

    class Config:
        json_schema_extra = {
            "example": {
                "username": "johndoe",
                "password": "securepassword123"
            },
            "description": "Schema for user login"
        }


class UserToken(BaseModel):
    """Schema for user token response"""
    access_token: str = Field(..., description="The access token")
    token_type: str = Field("bearer", description="The token type")

    class Config:
        json_schema_extra = {
            "example": {
                "access_token": "v4.local.example_token_content",
                "token_type": "bearer"
            }
        }