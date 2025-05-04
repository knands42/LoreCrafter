import os

from langchain_google_genai import ChatGoogleGenerativeAI, GoogleGenerativeAIEmbeddings
from langchain_openai import ChatOpenAI, OpenAIEmbeddings


class LLMFactory:
    @staticmethod
    def create_chat():
        if os.getenv("GOOGLE_API_KEY"):
            return ChatGoogleGenerativeAI(
                model="gemini-2.0-flash",
                temperature=0.8,
                api_key=os.getenv("GOOGLE_API_KEY")
            )
        elif os.getenv("OPENAI_API_KEY"):
            return ChatOpenAI(
                model="gpt-3.5-turbo",
                temperature=0.8,
                api_key=os.getenv("OPENAI_API_KEY")
            )
        else:
            raise ValueError("No API key found for chat. Please set GOOGLE_API_KEY or OPENAI_API_KEY in your environment.")

    @staticmethod
    def create_image_generator():
        if os.getenv("GOOGLE_API_KEY"):
            return ChatGoogleGenerativeAI(
                model="models/gemini-2.0-flash-exp-image-generation",
                api_key=os.getenv("GOOGLE_API_KEY")
            )
        elif os.getenv("OPENAI_API_KEY"):
            return ChatOpenAI(
                model="dall-e-3",
                api_key=os.getenv("OPENAI_API_KEY")
            )
        else:
            raise ValueError("No API key found for image generation. Please set GOOGLE_API_KEY or OPENAI_API_KEY in your environment.")

    @staticmethod
    def create_function_embeddings():
        if os.getenv("GOOGLE_API_KEY"):
            return GoogleGenerativeAIEmbeddings(
                model="models/embedding-001",
                api_key=os.getenv("GOOGLE_API_KEY")
            )
        elif os.getenv("OPENAI_API_KEY"):
            return OpenAIEmbeddings()
        else:
            raise ValueError("No API key found for embeddings. Please set GOOGLE_API_KEY or OPENAI_API_KEY in your environment.")
