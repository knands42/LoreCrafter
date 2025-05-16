import os

from peewee import PostgresqlDatabase


def get_db():
    db = PostgresqlDatabase(
        os.getenv("DATABASE_NAME", "lorecrafter"),
        user=os.getenv("DATABASE_USER", "postgres"),
        password=os.getenv("DATABASE_PASSWORD", "postgres"),
        host=os.getenv("DATABASE_HOST", "localhost"),
        port=os.getenv("DATABASE_PORT", "5432"),
    )

    db.connect()
    return db
