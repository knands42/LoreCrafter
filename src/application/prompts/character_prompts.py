from langchain.prompts import PromptTemplate
from langchain_core.messages import HumanMessage


def create_appearance_prompt(custom_appearance=None) -> PromptTemplate:
    if custom_appearance:
        return PromptTemplate.from_template("""
        Enhance the following appearance description of a TTRPG character with vivid detail and unique features.

        Race: {race}
        Gender: {gender}
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

        Race: {race}
        Gender: {gender}
        Story Tone: {tone}

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
        Gender: {gender}
        Race: {race}
        Alignment: {alignment}
        Appearance: {appearance}
        Story Tone: {tone}

        Existing Personality Description:
        {personality}

        Guidelines:
        - Preserve the core traits provided.
        - Add emotional depth, contradictions, or quirks that make the character feel real.
        - Describe motivations, behaviors, and notable personality traits in 1-2 sentences.
        - Reflect the character's alignment in their personality traits and moral compass.
        - Avoid clichés and make it feel grounded in a fantasy setting.
        - If relevant, subtly reflect how their appearance may influence behavior or social perception (only if it feels natural — this is rare).
        - Do not describe the character's alignment directly, make it narratively.

        Respond only with the enhanced personality.
        """)
    else:
        return PromptTemplate.from_template("""
        Create a deep and vivid personality for the following TTRPG character.

        Name: {name}
        Gender: {gender}
        Race: {race}
        Alignment: {alignment}
        Appearance: {appearance}
        Story Tone: {tone}

        Guidelines:
        - Write 1-2 sentences describing key traits, motivations, behaviors, flaws, and emotional tendencies.
        - Reflect the character's alignment in their personality traits and moral compass.
        - If relevant, subtly reflect how their appearance may influence behavior or social perception (only if it feels natural — this is rare).
        - Make it immersive and unique — suitable for storytelling or roleplaying.
        - Avoid listing traits — describe them narratively.
        - Do not describe the character's alignment directly, make it narratively.

        Respond only with the personality description.
        """)


def create_backstory_prompt(premade_story=None) -> PromptTemplate:
    base_format = """
## Early Life
Describe their origins: birthplace, family situation, early influences, and how their race/culture shaped them.

## Formative Events
Highlight 2-3 major events (losses, victories, betrayals, discoveries) that defined their character and shaped their outlook.

## Current Motivations
Summarize what drives them now: goals, inner conflicts, fears, ambitions, or unresolved past matters.

(Write in a style fitting the world theme, universe, and story tone.)
"""

    if premade_story:
        prompt = f"""
Enhance the TTRPG character’s existing backstory. Use the context below to deepen their journey while staying true to established elements.

Name: {{name}}
Gender: {{gender}}
Race: {{race}}
Alignment: {{alignment}}
Personality: {{personality}}
Appearance: {{appearance}}
Universe: {{universe}}
World Theme: {{world_theme}}
Story Tone: {{tone}}

Existing Backstory:
{{backstory}}

Guidelines:
1. Preserve the original narrative's key elements and emotional beats.
2. Expand on unanswered questions, mysterious hints, or impactful moments.
3. Maintain consistency with the character’s personality, appearance, and universe.
4. Add one seamless paragraph that elevates the emotional depth and complexity.
5. Keep the language tone consistent with what’s provided.

Use the following structure:
{base_format}

Respond only with the enriched story.
"""
    else:
        prompt = f"""
Create a rich, immersive backstory for the following TTRPG character. Use the provided attributes to guide the storytelling.

Name: {{name}}
Gender: {{gender}}
Race: {{race}}
Alignment: {{alignment}}
Personality: {{personality}}
Appearance: {{appearance}}
Universe: {{universe}}
World Theme: {{world_theme}}
Story Tone: {{tone}}

Instructions:
- Write a compelling backstory covering:
  • Early life (origins, culture, early influences)
  • 2-3 key events that shaped their beliefs or abilities
  • Current motivations, fears, ambitions, or mysteries
- Reflect the character's personality and appearance.
- Match the language style to the world theme and story tone.

Use the following structure:
{base_format}

Respond only with the final backstory.
"""

    return PromptTemplate.from_template(prompt)


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


def get_character_image_prompt(character_appearance: str) -> HumanMessage:
    return HumanMessage(f"""
    Create a portrait image of a character with the following appearance:
    {character_appearance}

    The image should be a high-quality character portrait suitable for a TTRPG character sheet.

    Respond only with the final portrait image.
    """)
