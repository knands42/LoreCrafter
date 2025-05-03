import os

from langchain_google_genai import ChatGoogleGenerativeAI
from langchain_openai import ChatOpenAI


class LLMFactory:
    @staticmethod
    def create():
        if os.getenv("GOOGLE_API_KEY"):
            return ChatGoogleGenerativeAI(
                model="gemini-2.0-flash",
                temperature=0.8,
                api_key=os.getenv("GOOGLE_API_KEY")
            )
        return ChatOpenAI(
            model="gpt-3.5-turbo",
            temperature=0.8,
            api_key=os.getenv("OPENAI_API_KEY")
        )

    @staticmethod
    def create_image_generator():
        return ChatGoogleGenerativeAI(
            model="models/gemini-2.0-flash-exp-image-generation",
            api_key=os.getenv("GOOGLE_API_KEY")
        )
