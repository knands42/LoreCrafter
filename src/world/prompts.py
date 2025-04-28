from langchain.prompts import PromptTemplate
from langchain_core.messages import HumanMessage


def create_world_history_prompt(premade_history=None) -> PromptTemplate:
    if premade_history:
        return PromptTemplate.from_template("""
        Enhance the existing world history for a TTRPG campaign. Use the context below to deepen the lore while staying true to established elements.

        World Name: {name}
        Universe: {universe_prompt}
        World Theme: {world_theme_prompt}
        Story Tone: {tone_prompt}

        Existing World History:
        {custom_history}

        Guidelines:
        1. Preserve the original narrative's key elements and historical events.
        2. Expand on unanswered questions, mysterious hints, or impactful moments.
        3. Maintain consistency with the world's theme, universe, and tone.
        4. Add one or two seamless paragraphs that elevate the world's depth and complexity.
        5. Keep the language tone consistent with what's provided.
        6. The idea is to enhance a background story for the world where the campaign will take place, and not the campaign stoy itself

        Respond only with the enriched world history.
        """)
    else:
        return PromptTemplate.from_template("""
        Create a rich, immersive world history for a TTRPG campaign. Use the provided attributes to guide the storytelling.

        World Name: {name}
        Universe: {universe_prompt}
        World Theme: {world_theme_prompt}
        Story Tone: {tone_prompt}

        Instructions:
        - Write 2-3 compelling paragraphs that include:
            • The world's origins and major historical eras
            • Significant events that shaped the current state of the world
            • Major conflicts, alliances, or power shifts
            • Cultural or technological developments
        - Stay true to the theme and universe provided.
        - Reflect the tone in your language.
        - Create a history that offers adventure hooks and mysteries for players to explore.
        - The idea is to create a background story for the world where the campaign will take place, and not the campaign stoy itself

        Respond only with the final world history.
        """)


def create_timeline_prompt(premade_timeline=None) -> PromptTemplate:
    if premade_timeline:
        return PromptTemplate.from_template("""
        Enhance the existing timeline for a TTRPG world. Use the context below to deepen the chronology while staying true to established elements.

        World Name: {name}
        World History: {history}
        Universe: {universe_prompt}
        World Theme: {world_theme_prompt}
        Story Tone: {tone_prompt}

        Existing Timeline:
        {custom_timeline}

        Guidelines:
        1. Preserve the original timeline's key events and dates.
        2. Add 3-5 additional significant events at appropriate points in the timeline.
        3. Maintain consistency with the world's history, theme, and universe.
        4. Include a mix of major historical events, cultural shifts, and potential adventure hooks.
        5. Format as a chronological list with dates/eras and brief descriptions.
        6. This will be used later on when creating the campaign that will be based on one of these timelines (or a mix) of these timelines

        Respond only with the enhanced timeline.
        """)
    else:
        return PromptTemplate.from_template("""
        Create a detailed timeline for a TTRPG world. Use the provided attributes to guide the chronology.

        World Name: {name}
        World History: {history}
        Universe: {universe_prompt}
        World Theme: {world_theme_prompt}
        Story Tone: {tone_prompt}

        Instructions:
        - Create a chronological timeline with 8-12 significant events
        - Include a mix of:
            • Ancient/founding events
            • Major conflicts or wars
            • Cultural or technological revolutions
            • Recent events that set the stage for adventures
        - Format each entry with a date/era and a brief description
        - Ensure consistency with the world history provided
        - Include events that could serve as adventure hooks or campaign backgrounds
        - This will be used later on when creating the campaign that will be based on one of these timelines (or a mix) of these timelines

        Respond only with the final timeline.
        """)


def create_campaign_setting_prompt(premade_campaign=None) -> PromptTemplate:
    if premade_campaign:
        return PromptTemplate.from_template("""
        Enhance the existing campaign setting for a TTRPG world. Use the context below to deepen the setting while staying true to established elements.

        World Name: {name}
        World History: {history}
        Timeline: {timeline}
        Universe: {universe_prompt}
        World Theme: {world_theme_prompt}
        Story Tone: {tone_prompt}

        Existing Campaign Setting:
        {custom_campaign}

        Guidelines:
        1. Preserve the original setting's key locations, factions, and conflicts.
        2. Expand on the current state of the world and immediate tensions.
        3. Add 2-3 additional locations, factions, or conflicts that enrich the setting.
        4. Maintain consistency with the world's history, timeline, theme, and universe.
        5. Include potential adventure hooks and campaign arcs.

        Respond only with the enhanced campaign setting.
        """)
    else:
        return PromptTemplate.from_template("""
        Create a detailed campaign setting for a TTRPG world. Use the provided attributes to guide the setting creation.

        World Name: {name}
        World History: {history}
        Timeline: {timeline}
        Universe: {universe_prompt}
        World Theme: {world_theme_prompt}
        Story Tone: {tone_prompt}

        Instructions:
        - Write 3-4 paragraphs describing the current state of the world as a campaign setting
        - Include:
            • 3-5 key locations where adventures might take place
            • 2-4 major factions or groups and their motivations
            • Current conflicts or tensions that drive the narrative
            • Unique features of the setting (magic systems, technology, cultural practices)
        - Ensure consistency with the world history and timeline provided
        - Create a setting rich with adventure hooks and storytelling opportunities

        Respond only with the final campaign setting.
        """)


def create_hidden_elements_prompt() -> PromptTemplate:
    return PromptTemplate.from_template("""
    Create a list of hidden elements and secrets for a TTRPG world. These should be unknown to most inhabitants but discoverable by players during campaigns.

    World Name: {name}
    World History: {history}
    Timeline: {timeline}
    Campaign Setting: {campaign}
    Universe: {universe_prompt}
    World Theme: {world_theme_prompt}
    Story Tone: {tone_prompt}

    Instructions:
    - Create 5-8 hidden elements or secrets that could be revealed during gameplay
    - Include a mix of:
        • Ancient mysteries or forgotten knowledge
        • Hidden locations or dungeons
        • Secret organizations or cults
        • Concealed magical artifacts or technology
        • Conspiracies or plots by powerful entities
    - Each secret should connect to the established world history or setting
    - Format as a list with brief descriptions (2-3 sentences each)
    - Include potential clues or ways players might discover these secrets

    Respond only with the final list of hidden elements.
    """)


def get_world_image_prompt(world_description: str) -> HumanMessage:
    return HumanMessage(f"""
    Create a landscape image representing a fantasy world with the following description:
    {world_description}

    The image should be a high-quality landscape suitable for a TTRPG campaign setting document.
    """)