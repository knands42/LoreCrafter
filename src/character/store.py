import json
import uuid

from dotenv import load_dotenv
from langchain.docstore.document import Document
from langchain_chroma import Chroma
from langchain_openai import OpenAIEmbeddings

load_dotenv()

embeddings = OpenAIEmbeddings()

vectorstore = Chroma(
    collection_name="characters",
    embedding_function=embeddings,
    persist_directory="chroma_db"
)


def store_character_in_chroma(created_character: dict[str, str]):
    doc = Document(page_content=json.dumps(created_character),
                   metadata={"id": str(uuid.uuid4()), "name": created_character['name']})

    vectorstore.add_documents([doc])


def search_similar_characters(query: str, top_k: int = 3):
    return vectorstore.similarity_search(query, k=top_k)
