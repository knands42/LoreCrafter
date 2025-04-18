import uuid

from langchain.docstore.document import Document
from langchain.embeddings import OpenAIEmbeddings
from langchain_chroma import Chroma
from langchain_openai import OpenAIEmbeddings


def store_character_in_chroma(created_character: dict[str, str]):
    embeddings = OpenAIEmbeddings()

    doc = Document(page_content=created_character.__str__(),
                   metadata={"id": str(uuid.uuid4()), "name": created_character['name']})

    vectorstore = Chroma(
        collection_name="characters",
        embedding_function=embeddings,
        persist_directory="chroma_db"
    )

    vectorstore.add_documents([doc])

