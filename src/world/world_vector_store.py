import json
from typing import List, Any
from uuid import uuid4

from langchain.docstore.document import Document
from langchain_chroma import Chroma
from langchain_openai import OpenAIEmbeddings


class WorldVectorStore:
    def __init__(
        self,
        collection_name: str = "worlds",
        persist_directory: str = "chroma_db",
        embedding_function: Any = None,
     ):
        self.embedding_function = embedding_function or OpenAIEmbeddings()
        self.vector_store = Chroma(
            collection_name=collection_name,
            embedding_function=self.embedding_function,
            persist_directory=persist_directory
        )

    def store(self, world_info: dict) -> None:
        """Store world information in the vector database.

        Args:
            world_info: A dictionary containing world information.
        """
        world_json = json.dumps(world_info)

        document = Document(
            page_content=world_json,
            metadata={
                "id": str(uuid4()),
                "name": world_info.get("name", "unknown")
            }
        )

        self.vector_store.add_documents([document])

    def search_similar(self, query: str, top_k: int = 1) -> List[Document]:
        """Search for similar worlds in the vector database.

        Args:
            query: The search query.
            top_k: The number of results to return.

        Returns:
            A list of documents containing similar worlds.
        """
        return self.vector_store.similarity_search(query, k=top_k)

    def get_by_id(self, world_id: str) -> dict | None:
        """Get a world by its ID.

        Args:
            world_id: The ID of the world to retrieve.

        Returns:
            The world information as a dictionary.
        """
        results = self.vector_store.similarity_search(
            f"id: {world_id}",
            k=1
        )

        if results:
            return json.loads(results[0].page_content)
        return None