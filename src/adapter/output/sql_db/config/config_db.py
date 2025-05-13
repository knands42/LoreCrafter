from psycopg2 import connect
import os


def get_db():
    db_url = os.getenv('DATABASE_URL', 'postgresql://localhost/lorecrafter')
    return connect(db_url)

def init_database(db, models):
    db.connect()
    db.create_tables(models)
    db.close()
