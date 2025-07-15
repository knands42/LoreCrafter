package dnd_5e

import (
	"fmt"
)

type DND5ECreateCampaignPromptData struct {
	WorldTheme    string
	WrittenTone   string
	CustomSetting string
}

func NewDND5ECreateCampaignPromptData(worldTheme, writtenTone, customSetting string) *DND5ECreateCampaignPromptData {
	return &DND5ECreateCampaignPromptData{
		WorldTheme:    worldTheme,
		WrittenTone:   writtenTone,
		CustomSetting: customSetting,
	}
}

func (prompt *DND5ECreateCampaignPromptData) CreateCampaignSettingPrompt() string {
	baseFormat := createCampaignSettingsMdFormat()

	if prompt.CustomSetting != "" {
		return fmt.Sprintf(`
Enhance the existing campaign setting for a TTRPG world. Use the context below to deepen the setting while staying true to established elements.

World Theme Setting: %s
Narrative Written Tone: %s

Existing Campaign Setting:
%s

Guidelines:
1. Preserve the original setting's key locations, factions, and conflicts.
2. Expand on the current state of the world and immediate tensions.
3. Add 2-3 additional locations, factions, or conflicts that enrich the setting.
4. Maintain consistency with the world's history, timeline, theme, and universe.
5. Include potential adventure hooks and campaign arcs.

Use the following format for your response:
%s

Respond only with the enhanced campaign setting, following the structure above.
`,
			prompt.WorldTheme,
			prompt.WrittenTone,
			prompt.CustomSetting,
			baseFormat,
		)
	} else {
		return fmt.Sprintf(`
Create a detailed campaign setting for a TTRPG world using the provided attributes.

World Theme Setting: %s
Narrative Written Tone: %s

Instructions:
- Write using the format below.
- Ensure consistency with the provided world history, timeline, universe, and theme.
- Richly describe locations, factions, and conflicts to inspire storytelling.
- Create a setting full of potential adventure hooks and campaign arcs.

Use the following format for your response:
%s

Respond only with the final campaign setting, following the structure above.
`,
			prompt.WorldTheme,
			prompt.WrittenTone,
			baseFormat,
		)
	}
}

func createCampaignSettingsMdFormat() string {
	return `
## Introduction
(Brief overview introducing the world, its theme, and its tone.)

---

## World Background
- **World History:** (Key points from {backstory})
- **Timeline Highlights:** (Notable events that happened before and the current state of the world)
- **Universe Overview:** (Summary from the universe where the world takes place (if relevant))
- **World Theme:** (Highlights the theme of the world, it may not be the entire world's that follow the same theme settings, but the main adventure should highlight that feeling)

---

## Key Locations (that can be relevant to the campaign directly or undirectly)
**Location 1**  
(2-4 sentence description)

**Location 2**  
(2-4 sentence description)

**Location 3**  
(2-4 sentence description)

---

## Major Factions / Organizations / Clans (what makes the most sense)
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

## Hidden Elements
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

---

## Adventure Hooks
- **Hook 1:** (1–2 sentence idea)
- **Hook 2:** (1–2 sentence idea)
- **Hook 3:** (1–2 sentence idea)
`
}
