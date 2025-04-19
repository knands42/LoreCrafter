import json
from typing import Any
from uuid import uuid4

from langchain.docstore.document import Document
from langchain_chroma import Chroma
from langchain_openai import OpenAIEmbeddings


class CharacterVectorStore:
    def __init__(
        self,
        collection_name: str = "characters",
        persist_directory: str = "chroma_db",
        embedding_function: Any = None,
    ):
        self.embedding_function = embedding_function or OpenAIEmbeddings()
        self.vectorstore = Chroma(
            collection_name=collection_name,
            embedding_function=self.embedding_function,
            persist_directory=persist_directory,
        )
        self.retriever = self.vectorstore.as_retriever()

    def store(self, character: dict[str, str]) -> None:
        doc = Document(
            page_content=json.dumps(character),
            metadata={
                "id": str(uuid4()),
                "name": character.get("name", "unknown")
            }
        )
        self.vectorstore.add_documents([doc])

    def search_similar(self, query: str, top_k: int = 3) -> list[Document]:
        return self.vectorstore.similarity_search(query, k=top_k)

    def retrieve_relevant_docs(self, query: str) -> list[Document]:
        return self.retriever.get_relevant_documents(query)
