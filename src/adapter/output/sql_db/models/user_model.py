from datetime import datetime, UTC

from peewee import Model, CharField, BooleanField, DateTimeField

from src.adapter.output.sql_db.config.config import get_db


class UserModel(Model):
    username = CharField(unique=True, max_length=80)
    email = CharField(unique=True, max_length=120)
    password_hash = CharField()
    is_active = BooleanField(default=True)
    created_at = DateTimeField(default=datetime.now(UTC))
    updated_at = DateTimeField(default=datetime.now(UTC))

    class Meta:
        database = get_db()

