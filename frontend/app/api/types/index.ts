// Character types
export interface Character {
  id: string;
  name: string;
  gender: string;
  race: string;
  personality: string;
  appearance: string;
  universe: string;
  world_theme: string;
  tone: string;
  backstory: string;
  custom_story?: string;
  linked_world_id?: string;
  image_filename?: string;
}

export interface CreateCharacterRequest {
  name: string;
  gender: string;
  race: string;
  personality?: string;
  appearance?: string;
  universe: string;
  world_theme: string;
  tone: string;
  custom_story?: string;
  linked_world_id?: string;
}

// World types
export interface World {
  id: string;
  name: string;
  universe: string;
  world_theme: string;
  tone: string;
  history: string;
  timeline: string;
  custom_history?: string;
  custom_timeline?: string;
  custom_backstory?: string;
  image_filename?: string;
}

export interface CreateWorldRequest {
  name: string;
  universe: string;
  world_theme: string;
  tone: string;
  custom_history?: string;
  custom_timeline?: string;
  custom_backstory?: string;
}

// Campaign types
export interface Campaign {
  id: string;
  name: string;
  universe: string;
  world_theme: string;
  tone: string;
  campaign: string;
  hidden_elements: string;
  custom_campaign?: string;
  linked_world_id?: string;
  linked_character_ids?: string[];
}

export interface CreateCampaignRequest {
  name: string;
  universe: string;
  world_theme: string;
  tone: string;
  custom_campaign?: string;
  linked_world_id?: string;
  linked_character_ids?: string[];
}

// Common types for dropdown options
export interface SelectOption {
  value: string;
  label: string;
}

// Universe options
export const universeOptions: SelectOption[] = [
  { value: 'D&D', label: 'Dungeons & Dragons' },
  { value: 'Pathfinder', label: 'Pathfinder' },
  { value: 'Warhammer', label: 'Warhammer' },
  { value: 'Call of Cthulhu', label: 'Call of Cthulhu' },
  { value: 'Vampire: The Masquerade', label: 'Vampire: The Masquerade' },
  { value: 'Custom', label: 'Custom Universe' },
];

// World theme options
export const worldThemeOptions: SelectOption[] = [
  { value: 'fantasy', label: 'Fantasy' },
  { value: 'sci-fi', label: 'Science Fiction' },
  { value: 'steampunk', label: 'Steampunk' },
  { value: 'post-apocalyptic', label: 'Post-Apocalyptic' },
  { value: 'horror', label: 'Horror' },
  { value: 'cyberpunk', label: 'Cyberpunk' },
  { value: 'historical', label: 'Historical' },
];

// Tone options
export const toneOptions: SelectOption[] = [
  { value: 'Epic', label: 'Epic' },
  { value: 'Gritty', label: 'Gritty' },
  { value: 'Humorous', label: 'Humorous' },
  { value: 'Dark', label: 'Dark' },
  { value: 'Whimsical', label: 'Whimsical' },
  { value: 'Mysterious', label: 'Mysterious' },
];