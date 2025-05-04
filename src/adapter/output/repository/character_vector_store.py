import json
from typing import Any

from langchain.docstore.document import Document
from langchain_chroma import Chroma

from src.adapter.output.llm import LLMFactory
from src.adapter.output.utils import UUIDEncoder
from src.application.domain.character_domain import CharacterDomain


class CharacterVectorStore:
    def __init__(
        self,
        collection_name: str = "characters",
        persist_directory: str = "chroma_db",
        embedding_function: Any = None,
    ):
        self.embedding_function = embedding_function if embedding_function else LLMFactory.create_function_embeddings()
        self.vectorstore = Chroma(
            collection_name=collection_name,
            embedding_function=self.embedding_function,
            persist_directory=persist_directory,
        )
        self.retriever = self.vectorstore.as_retriever()

    def store(self, character: CharacterDomain) -> None:
        doc = Document(
            page_content=json.dumps(character, cls=UUIDEncoder),
            metadata={
                "id": str(character['id']),
                "name": character.get("name") or "unknown",
                "world_linked": character.get("linked_world_id") or "unknown",
            }
        )
        
        self.vectorstore.add_documents([doc])

    def search_similar(self, query: str, top_k: int = 3) -> list[Document]:
        return self.vectorstore.similarity_search(query, k=top_k)

    def retrieve_relevant_docs(self, query: str) -> list[Document]:
        return self.retriever.get_relevant_documents(query)
