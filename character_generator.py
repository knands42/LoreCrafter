from langchain_openai import ChatOpenAI
from langchain.prompts import PromptTemplate
from langchain.schema.output_parser import StrOutputParser
from dotenv import load_dotenv
import os

load_dotenv()

llm = ChatOpenAI(
    model="gpt-3.5-turbo", 
    temperature=0.8,
    openai_api_key=os.getenv("OPENAI_API_KEY"),
    api_key=os.getenv("OPENAI_API_KEY")
)

def generate_character(character_info):
    prompt = PromptTemplate.from_template("""
    Create a detailed backstory for a fantasy RPG character with the following information:
    
    Name: {name}
    Race: {race}
    Personality: {personality}
    
    Generate a compelling 2-3 paragraph backstory that explains:
    - Where they came from
    - What shaped their personality
    - What motivates them
    - Any significant events in their past
    
    Make the story rich in detail and consistent with fantasy RPG tropes.
    """)
    
    chain = prompt | llm | StrOutputParser()
    result = chain.invoke(character_info)
    return result