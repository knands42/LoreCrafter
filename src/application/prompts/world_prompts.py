from langchain.prompts import PromptTemplate
from langchain_core.messages import HumanMessage


def create_world_history_prompt(premade_history=None) -> PromptTemplate:
    base_format = """
## Origins
Describe how the world was created or how its earliest civilizations began. Mention any gods, cosmic events, or mythologies if relevant, keep it simple if non relevant.

## Major Historical Eras
Break the world's history into 2-3 key eras (e.g., Age of Kings, Time of Chaos, Age of Renewal) with a brief overview of each.

## Key Events
Highlight 3-5 pivotal moments (wars, alliances, cataclysms, discoveries) that shaped the world into its current state.

## Current Setting
Summarize the world's current status: tensions, mysteries, and potential adventure seeds.

(Write in a style fitting the provided world theme and tone.)
"""

    if premade_history:
        prompt = f"""
Enhance the existing world history for a TTRPG campaign. Use the context below to deepen the lore while staying true to established elements.

World Name: {{name}}
Universe: {{universe_prompt}}
World Theme: {{world_theme_prompt}}
Story Tone: {{tone_prompt}}

Existing World History:
{{custom_history}}

Guidelines:
1. Preserve the original narrative's key elements and historical events.
2. Expand on unanswered questions, mysterious hints, or impactful moments.
3. Maintain consistency with the world's theme, universe, and tone.
4. Add 1-2 short but immersive paragraphs to elevate the world's depth and complexity.
5. Keep the language tone consistent with what's provided.
6. Focus on the world’s background story — not the campaign story itself.

Use the following structure:
{base_format}

Respond only with the enriched world history.
"""
    else:
        prompt = f"""
Create a rich, immersive world history for a TTRPG campaign. Use the provided attributes to guide the storytelling.

World Name: {{name}}
Universe: {{universe_prompt}}
World Theme: {{world_theme_prompt}}
Story Tone: {{tone_prompt}}

Instructions:
- Write 2-3 compelling paragraphs covering:
  • The world's origins and its earliest civilizations.
  • Major historical eras and defining events.
  • Cultural, magical, or technological developments.
  • Conflicts, alliances, or power shifts.
  • Mysteries or open questions that could lead to adventures.
- Reflect the provided world theme and story tone.
- Focus only on background lore — not on specific campaign stories.

Use the following structure:
{base_format}

Respond only with the final world history.
"""

    return PromptTemplate.from_template(prompt)


def create_timeline_prompt(premade_timeline=None) -> PromptTemplate:
    base_format = """
## Events

**[Era/Year]** — **Event Title**  
Short description of the event. (1-3 sentences.)

**[Era/Year]** — **Event Title**  
Short description of the event. (1-3 sentences.)

**[Era/Year]** — **Event Title**  
Short description of the event. (1-3 sentences.)

(Repeat as needed for 8-12 entries.)

"""

    if premade_timeline:
        prompt = f"""
Enhance the existing timeline for a TTRPG world. Use the context below to deepen the chronology while staying true to established elements.

World Name: {{name}}
World History: {{history}}
Universe: {{universe_prompt}}
World Theme: {{world_theme_prompt}}
Story Tone: {{tone_prompt}}

Existing Timeline:
{{custom_timeline}}

Guidelines:
1. Preserve the original timeline's key events and dates.
2. Add 3-5 additional significant events at appropriate points in the timeline.
3. Maintain consistency with the world's history, theme, and universe.
4. Include a mix of major historical events, cultural shifts, and potential adventure hooks.
5. Format the timeline using the structure below: one line for era/year + event title, followed by a short paragraph for the description.
6. Add a small optional section for "Notes for Adventure Hooks" if needed.
7. This will be used later when creating campaigns based on these timelines.

Use the following format:
{base_format}

Respond only with the enhanced timeline.
"""
    else:
        prompt = f"""
Create a detailed timeline for a TTRPG world. Use the provided attributes to guide the chronology.

World Name: {{name}}
World History: {{history}}
Universe: {{universe_prompt}}
World Theme: {{world_theme_prompt}}
Story Tone: {{tone_prompt}}

Instructions:
- Create a chronological timeline with 8-12 significant events.
- Include a mix of:
  • Ancient/founding events
  • Major conflicts or wars
  • Cultural or technological revolutions
  • Recent events that set the stage for adventures
- Each event should have:
  • A date/era and title (bold)
  • A short paragraph (1-3 sentences) explaining the event
- Ensure consistency with the world history provided.
- Optionally highlight any events that could be used as adventure hooks.
- This will be used later when creating campaigns based on these timelines.

Use the following format:
{base_format}

Respond only with the final timeline.
"""

    return PromptTemplate.from_template(prompt)


def get_world_image_prompt(world_description: str) -> HumanMessage:
    return HumanMessage(f"""
    Create a landscape image representing a fantasy world with the following description:
    {world_description}

    The image should be a high-quality landscape suitable for a TTRPG campaign setting document.
    """)
