# Dependencies
from typing import Optional

from adapter.output.sql_db.models.user_model import UserModel
from fastapi import APIRouter, Depends, HTTPException, Response, Cookie
from fastapi.security import OAuth2PasswordRequestForm

from src.adapter.output.sql_db.config.config_db import get_db, initialize_db
from src.application.domain.user_domain import UserCreation as UserCreationSchema, User, UserToken
from src.application.usecases.user_creation import UserCreation

# Create router
router = APIRouter()


def get_sql_db():
    db = get_db()
    initialize_db(UserModel)
    return db


def get_user_creation(db=Depends(get_sql_db)):
    return UserCreation(db)


# API endpoints
@router.post("/users", response_model=User)
async def create_user(
    user_data: UserCreationSchema,
    user_creation: UserCreation = Depends(get_user_creation)
):
    """
    Create a new user.

    This endpoint creates a new user with the provided username, email, and password.
    The password is hashed using argon2 before being stored in the database.
    """
    try:
        result = user_creation.create_user(user_data)
        if not result:
            raise HTTPException(status_code=400, detail="User already exists")

        return result
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error creating user: {str(e)}")


@router.post("/signin", response_model=UserToken)
async def signin(
    response: Response,
    form_data: OAuth2PasswordRequestForm = Depends(),
    user_creation: UserCreation = Depends(get_user_creation)
):
    """
    Sign in a user and return a token.

    This endpoint authenticates a user with the provided username and password,
    and returns a Paseto token that can be used for authentication in subsequent requests.
    The token is also set as a cookie.
    """
    try:
        # Get the user
        user = user_creation.get_user_by_username(form_data.username)
        if not user:
            raise HTTPException(status_code=401, detail="Invalid username or password")

        # Verify the password
        from src.adapter.output.sql_db.models.user_model import UserModel
        db_user = UserModel.get(UserModel.username == form_data.username)
        if not user_creation.verify_password(db_user.password_hash, form_data.password):
            raise HTTPException(status_code=401, detail="Invalid username or password")

        # Generate a token
        token = user_creation.generate_token(str(user.id))

        # Set the token as a cookie
        response.set_cookie(
            key="access_token",
            value=token,
            httponly=True,
            secure=True,  # Set to False in development
            samesite="lax",
            max_age=86400  # 1 day
        )

        # Return the token
        return UserToken(access_token=token, token_type="bearer")
    except HTTPException:
        raise
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error signing in: {str(e)}")


@router.get("/me", response_model=User)
async def get_current_user(
    access_token: Optional[str] = Cookie(None),
    user_creation: UserCreation = Depends(get_user_creation)
):
    """
    Get the current user.

    This endpoint returns the current user based on the access_token cookie.
    """
    if not access_token:
        raise HTTPException(status_code=401, detail="Not authenticated")

    try:
        # Verify the token
        payload = user_creation.verify_token(access_token)
        user_id = payload.get("sub")

        # Get the user from the database
        from src.adapter.output.sql_db.models.user_model import UserModel
        try:
            user = UserModel.get(UserModel.id == user_id)
            return User(
                id=user.id,
                username=user.username,
                email=user.email,
                is_active=user.is_active,
                created_at=user.created_at,
                updated_at=user.updated_at
            )
        except UserModel.DoesNotExist:
            raise HTTPException(status_code=404, detail="User not found")
    except ValueError as e:
        raise HTTPException(status_code=401, detail=str(e))
    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Error getting current user: {str(e)}")
