from langchain.prompts import PromptTemplate


def create_campaign_setting_prompt(premade_setting=None) -> PromptTemplate:
    base_format = """
## Introduction
(Brief overview introducing the world, its theme, and its tone.)

---

## World Background
- **World History:** (Key points from {history})
- **Timeline Highlights:** (Notable events from {timeline})
- **Universe Overview:** (Summary from {universe_prompt})
- **World Theme:** (Summary from {world_theme_prompt})
- **Story Tone:** (Summary from {tone_prompt})

---

## Key Locations
**Location 1**  
(2-4 sentence description)

**Location 2**  
(2-4 sentence description)

**Location 3**  
(2-4 sentence description)

---

## Major Factions
- **Faction 1:** (Summary: motivations, rivals, influence)
- **Faction 2:** (Summary: motivations, rivals, influence)

---

## Current Conflicts
- Conflict 1: (Brief description)
- Conflict 2: (Brief description)

---

## Unique Features
- **Magic/Technology Systems:** (Brief description)
- **Cultural Practices:** (Brief description)

---

## Adventure Hooks
- **Hook 1:** (1–2 sentence idea)
- **Hook 2:** (1–2 sentence idea)
- **Hook 3:** (1–2 sentence idea)
"""

    if premade_setting:
        prompt = f"""
Enhance the existing campaign setting for a TTRPG world. Use the context below to deepen the setting while staying true to established elements.

World Name: {{name}}
World History: {{history}}
Timeline: {{timeline}}
Universe: {{universe_prompt}}
World Theme: {{world_theme_prompt}}
Story Tone: {{tone_prompt}}

Existing Campaign Setting:
{{custom_setting}}

Guidelines:
1. Preserve the original setting's key locations, factions, and conflicts.
2. Expand on the current state of the world and immediate tensions.
3. Add 2-3 additional locations, factions, or conflicts that enrich the setting.
4. Maintain consistency with the world's history, timeline, theme, and universe.
5. Include potential adventure hooks and campaign arcs.

Use the following format for your response:
{base_format}

Respond only with the enhanced campaign setting, following the structure above.
"""
    else:
        prompt = f"""
Create a detailed campaign setting for a TTRPG world using the provided attributes.

World Name: {{name}}
World History: {{history}}
Timeline: {{timeline}}
Universe: {{universe_prompt}}
World Theme: {{world_theme_prompt}}
Story Tone: {{tone_prompt}}

Instructions:
- Write using the format below.
- Ensure consistency with the provided world history, timeline, universe, and theme.
- Richly describe locations, factions, and conflicts to inspire storytelling.
- Create a setting full of potential adventure hooks and campaign arcs.

Use the following format for your response:
{base_format}

Respond only with the final campaign setting, following the structure above.
"""

    return PromptTemplate.from_template(prompt)


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
