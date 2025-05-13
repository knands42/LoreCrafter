from psycopg2 import connect
import os

from src.adapter.output.sql_db.models.user_model import UserModel


def get_db():
    db_url = os.getenv('DATABASE_URL', 'postgresql://localhost/lorecrafter')
    return connect(db_url)

def init_database(db):
    db.connect()
    db.create_tables([UserModel])
    db.close()
