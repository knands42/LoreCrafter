from langchain.prompts import PromptTemplate


def create_appearance_prompt(appearance=None) -> PromptTemplate:
    if appearance:
        return PromptTemplate.from_template("""
        Enhance the following appearance description of a TTRPG character with vivid detail and unique features.

        Name: {name}
        Race: {race}
        Story Tone: {tone}

        Existing Appearance Description:
        {appearance}

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
        Story Tone: {tone}

        Guidelines:
        - Write 1–2 vivid sentences describing facial features, hair, build, clothing, and any unique characteristics.
        - Infuse light fantasy elements suitable for a TTRPG setting.
        - Use immersive language, avoiding lists or plain adjectives.

        Respond only with the appearance description.
        """)


def create_personality_prompt(personality=None) -> PromptTemplate:
    if personality:
        return PromptTemplate.from_template("""
        Enhance the personality description of the following TTRPG character with depth, nuance, and immersion.

        Name: {name}
        Race: {race}
        Appearance: {appearance}
        Story Tone: {tone}

        Existing Personality Description:
        {personality}

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
        Story Tone: {tone}

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
        World Theme: {world_theme}
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
        World Theme: {world_theme}
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


def summarize_backstory_prompt(backstory: str) -> PromptTemplate:
    return PromptTemplate.from_template(f"""
    Summarize the character's background story below.
    
    Background Story:
    {backstory}
    
    Instructions:
    - Summarize the backstory concisely.
    - Limit the response to one or two sentences at most.
    - Focus on the most relevant aspects of the character's past that define their identity or motivation.
    - Do not add any extra information or interpretation.
    
    Respond only with the final summarized backstory.
    """)
