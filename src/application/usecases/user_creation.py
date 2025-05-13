import json
import os
import secrets
from datetime import datetime, timedelta, UTC
from typing import Optional

import pyseto
from pyargon2 import hash, Argon2Error
from pyseto import Key
from rich import print

from src.adapter.output.sql_db.models.user_model import UserModel
from src.application.domain.user_domain import User, UserCreation as UserCreationSchema


class UserCreation:
    def __init__(self, db):
        self.db = db
        self.salt = os.urandom(16).hex()

        self.__init_paseto_keys()

    def __init_paseto_keys(self):
        """Initialize Paseto keys for token generation and verification."""
        # TODO: Load keys from a secure secret management system
        private_key_pem = open(os.environ.get("PASETO_PRIVATE_KEY"), "rb").read()
        public_key_pem = open(os.environ.get("PASETO_PUBLIC_KEY"), "rb").read()
        self.__private_key = Key.new(version=4, purpose="public", key=private_key_pem)
        self.__public_key = Key.new(version=4, purpose="public", key=public_key_pem)

    def create_user(self, user_data: UserCreationSchema) -> User:
        """Create a new user with hashed password."""
        password_hash = hash(user_data.password, self.salt)

        user = UserModel.create(
            username=user_data.username,
            email=user_data.email,
            password_hash=password_hash,
            is_active=True,
            created_at=datetime.now(UTC),
            updated_at=datetime.now(UTC)
        )

        return User(
            id=user.id,
            username=user.username,
            email=user.email,
            is_active=user.is_active,
            created_at=user.created_at,
            updated_at=user.updated_at
        )

    def verify_password(self, password_hash: str, password: str) -> bool:
        """Verify a password against a hash."""
        try:
            return hash(password, self.salt) == password_hash
        except Argon2Error:
            print("Error verifying password")
            return False

    def generate_token(self, user_info: User, expires_in: timedelta = timedelta(days=1)) -> bytes:
        """Generate a Paseto token for a user."""
        now = datetime.now(UTC)
        expiration = now + expires_in
        token_id = secrets.token_hex(16)

        # Create the token payload
        payload = {
            'sub': str(user_info.id),
            'username': user_info.username,
            'email': user_info.email,
            'is_active': user_info.is_active,
            'iat': now.timestamp(),
            'exp': expiration.timestamp(),
            'jti': token_id
        }

        return pyseto.encode(
            self.__private_key,
            json.dumps(payload).encode('utf-8'),
            footer=json.dumps({"kid": self.__public_key.to_paserk_id()}).encode('utf-8'),
            serializer=json,
            exp=int(expiration.timestamp()),
        )

    def verify_token(self, token: bytes) -> dict:
        """Verify a Paseto token and return the payload."""
        try:
            token_decoded = pyseto.decode(
                self.__public_key,
                token,
                deserializer=json,
            )

            if token_decoded.payload.get('exp', 0) < datetime.now(UTC).timestamp():
                raise ValueError("Token has expired")

            return token_decoded.payload
        except Exception as e:
            raise ValueError(f"Invalid token: {str(e)}")

    def get_user_by_username(self, username: str) -> Optional[User]:
        """Get a user by username."""
        try:
            user = UserModel.get(UserModel.username == username)
            return User(
                id=user.id,
                username=user.username,
                email=user.email,
                is_active=user.is_active,
                created_at=user.created_at,
                updated_at=user.updated_at
            )
        except Exception as e:
            print(f"Error fetching user: {str(e)}")
            return None
