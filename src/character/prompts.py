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


def create_appearance_prompt(custom_appearance=None) -> PromptTemplate:
    if custom_appearance:
        return PromptTemplate.from_template("""
        Enhance the following appearance description of a TTRPG character with vivid detail and unique features.

        Name: {name}
        Race: {race}

        Existing Appearance Description:
        {custom_appearance}

        Guidelines:
        - Preserve the existing description while expanding on distinct physical traits, expressions, and presence.
        - Reflect fantasy elements subtly (clothing, posture, magical touches, etc.).
        - Write 1–2 immersive sentences, rich with sensory language.

        Respond only with the enhanced appearance.
        """)
    else:
        return PromptTemplate.from_template("""
        Describe the physical appearance of the following TTRPG character.

        Name: {name}
        Race: {race}

        Guidelines:
        - Write 1–2 vivid sentences describing facial features, hair, build, clothing, and any unique characteristics.
        - Infuse light fantasy elements suitable for a TTRPG setting.
        - Use immersive language, avoiding lists or plain adjectives.

        Respond only with the appearance description.
        """)


def create_personality_prompt(custom_personality=None) -> PromptTemplate:
    if custom_personality:
        return PromptTemplate.from_template("""
        Enhance the personality description of the following TTRPG character with depth, nuance, and immersion.

        Name: {name}
        Race: {race}
        Appearance: {appearance}

        Existing Personality Description:
        {custom_personality}

        Guidelines:
        - Preserve the core traits provided.
        - Add emotional depth, contradictions, or quirks that make the character feel real.
        - Describe motivations, behaviors, and notable personality traits in 1-2 sentences.
        - Avoid clichés and make it feel grounded in a fantasy setting.
        - If relevant, subtly reflect how their appearance may influence behavior or social perception (only if it feels natural — this is rare).

        Respond only with the enhanced personality.
        """)
    else:
        return PromptTemplate.from_template("""
        Create a deep and vivid personality for the following TTRPG character.

        Name: {name}
        Race: {race}
        Appearance: {appearance}

        Guidelines:
        - Write 1-2 sentences describing key traits, motivations, behaviors, flaws, and emotional tendencies.
        - If relevant, subtly reflect how their appearance may influence behavior or social perception (only if it feels natural — this is rare).
        - Make it immersive and unique — suitable for storytelling or roleplaying.
        - Avoid listing traits — describe them narratively.

        Respond only with the personality description.
        """)


def create_backstory_prompt(premade_story=None) -> PromptTemplate:
    if premade_story:
        return PromptTemplate.from_template("""
        Enhance the TTRPG character’s existing backstory. Use the context below to deepen their journey while staying true to established elements.

        Name: {name}
        Race: {race}
        Personality: {personality}
        Appearance: {appearance}
        Universe: {universe}
        Universe Theme: {universe_theme}
        Story Tone: {tone}

        Existing Backstory:
        {custom_story}

        Guidelines:
        1. Preserve the original narrative's key elements and emotional beats.
        2. Expand on unanswered questions, mysterious hints, or impactful moments.
        3. Maintain consistency with the character’s personality, appearance, and universe.
        4. Add one seamless paragraph that elevates the story's emotional depth and complexity.
        5. Keep the language tone consistent with what’s provided.

        Respond only with the enriched story.
        """)
    else:
        return PromptTemplate.from_template("""
        Create a rich, immersive backstory for the following TTRPG character. Use the provided attributes to guide the storytelling.

        Name: {name}
        Race: {race}
        Personality: {personality}
        Appearance: {appearance}
        Universe: {universe}
        Universe Theme: {universe_theme}
        Story Tone: {tone}

        Instructions:
        - Write one compelling paragraph that includes:
            • Their origins and upbringing
            • Events that shaped who they are
            • Their internal motivations and current goals
            • Any defining relationships or pivotal experiences
        - Stay true to the personality and appearance provided.
        - Reflect the tone and universe style in your language.

        Respond only with the final backstory paragraph.
        """)
