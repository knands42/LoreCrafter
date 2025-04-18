from langchain.prompts import PromptTemplate


def get_universe() -> dict[str, str]:
    """Get the context descriptions for different TTRPG universes."""
    return {
        "d&d": "In the D&D universe, adventurers explore a rich fantasy world filled with dragons, ancient magic, and medieval kingdoms full of mystery.",
        "call of cthulhu": "The Call of Cthulhu universe plunges characters into Lovecraftian horror, where cosmic secrets unravel sanity and danger lurks in the unknown.",
        "starfinder": "In the Starfinder universe, science fiction meets fantasy as players journey through galaxies filled with aliens, starships, and arcane technology.",
        "pathfinder": "Pathfinder offers a deep high-fantasy universe teeming with diverse cultures, sprawling lore, and epic quests driven by heroism and conflict.",
        "vampire the masquerade": "Vampire The Masquerade immerses players in a modern gothic world of vampires, where political intrigue and personal horror shape every dark corner."
    }


def get_universe_theme() -> dict[str, str]:
    """Get thematic keywords and their descriptions that represent different story genres and settings."""
    return {
        "cyberpunk": "In a cyberpunk world, advanced technology collides with societal decay, where neon-lit cities are ruled by powerful corporations.",
        "solarpunk": "A solarpunk setting envisions a hopeful future shaped by ecological balance, community resilience, and sustainable technology.",
        "steampunk": "The steampunk genre blends Victorian aesthetics with steam-powered inventions, crafting a world of gears, goggles, and bold exploration.",
        "apocalypse": "An apocalypse theme throws characters into a world ravaged by catastrophe, where survival is uncertain and civilization is a memory.",
        "noir": "In a noir setting, the story unfolds in a moody atmosphere of crime, moral ambiguity, and shadowy figures working in the dark.",
        "fantasy": "A fantasy universe is filled with magic, mythical creatures, and grand quests, where heroes rise to challenge fate and ancient forces.",
        "space opera": "Space opera delivers an epic journey across galaxies, rich with interstellar politics, starship battles, and legendary destinies.",
        "gothic horror": "Gothic horror evokes eerie dread and beauty, where crumbling castles, haunting secrets, and supernatural terror abound.",
        "mythology": "A mythology theme draws on ancient tales of gods, epic battles, and heroic quests passed down through generations.",
        "urban fantasy": "Urban fantasy brings the mystical into the modern day, where magic, monsters, and hidden realms exist just beneath the surface of everyday life.",
        "default": "The same theme as the main universe"
    }


def get_tone_context() -> dict[str, str]:
    """Get the context descriptions for different story tones."""
    return {
        "happy": "A happy tone that brings light and joy to the story, focusing on uplifting experiences, triumphs, and heartfelt connections.",
        "sad": "A sad tone that creates emotional depth, revealing stories of loss, introspection, and personal struggle that touch the heart.",
        "dark": "A dark tone that creates a somber atmosphere where characters face grim realities, moral ambiguity, and difficult choices.",
        "hopeful": "A hopeful tone that inspires perseverance, highlighting resilience, growth, and the belief that things can get better against all odds.",
        "mysterious": "A mysterious tone that weaves intrigue and suspense into the story, where secrets unfold slowly and the unknown keeps readers guessing.",
        "epic": "An epic tone that elevates the story to legendary proportions, with grand challenges, heroic deeds, and destinies that shape entire worlds."
    }


def create_personality_prompt() -> PromptTemplate:
    pass

def create_appearance_prompt () -> PromptTemplate:
    pass

def create_backstory_enhancement_prompt() -> PromptTemplate:
    """Create a prompt template for enhancing an existing character story."""
    return PromptTemplate.from_template("""
    Enhance the character’s story based on the following information. The story should align with the specified universe, tone, theme and character attributes.

    Name: {name}
    Race: {race}
    Personality: {personality}
    Universe: {universe}
    Universe Theme: {universe_theme}
    Story Tone: {tone}

    Current Story:
    {custom_story}

    Enhancement Guidelines:
    1. Keep the core elements of the original story intact.
    2. Enrich the story with vivid detail and emotional depth.
    3. Expand on intriguing elements or untapped moments.
    4. Ensure everything stays true to the character’s race, personality, and universe.
    5. Maintain the tone throughout as specified.
    6. Add one new paragraph of content that seamlessly integrates into the existing narrative.

    """)


def create_backstory_prompt() -> PromptTemplate:
    """Create a prompt template for generating a new character backstory."""
    return PromptTemplate.from_template("""
    Create a rich and immersive backstory for the character below. The story should align with the specified universe, tone, theme and character attributes.

    Name: {name}
    Race: {race}
    Personality: {personality}
    Universe: {universe}
    Universe Theme: {universe_theme}
    Story Tone: {tone}

    Backstory Requirements:
    Write one compelling paragraph that includes:
    - Their origins and background
    - Key events that shaped their personality
    - Their motivations and goals
    - Any defining moments or relationships

    The story should:
    - Reflect the character’s attributes and the universe’s style
    - Stay consistent with the tone and theme provided
    - Be immersive, descriptive, and emotionally engaging
    
    """)
