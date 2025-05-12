// types.ts

export type Gender = 'male' | 'female' | 'non-binary' | 'other';
export type Alignment =
  | 'Lawful Good'
  | 'Neutral Good'
  | 'Chaotic Good'
  | 'Lawful Neutral'
  | 'True Neutral'
  | 'Chaotic Neutral'
  | 'Lawful Evil'
  | 'Neutral Evil'
  | 'Chaotic Evil';

export interface CharacterRequest {
  name: string;
  gender: Gender;
  race: string;
  personality?: string;
  appearance?: string;
  universe: string;
  world_theme: string;
  tone: string;
  alignment: Alignment;
  custom_story?: string;
  linked_world_id?: string;
}

export interface CharacterResponse {
  id: string;
  name: string;
  gender: Gender;
  race: string;
  personality: string;
  appearance: string;
  universe: string;
  world_theme: string;
  tone: string;
  backstory: string;
  custom_story?: string;
  linked_world_id?: string;
  image_filename: string;
}

export interface CharacterSearchQuery {
  query: string;
  top?: number;
}

export type CharacterSearchResponse = CharacterResponse[];

export interface WorldRequest {
  name: string;
  universe: string;
  world_theme: string;
  tone: string;
  custom_history?: string;
  custom_timeline?: string;
  custom_backstory?: string;
}

export interface WorldResponse {
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
  image_filename: string;
}

export interface WorldSearchQuery {
  query: string;
  top?: number;
}

export type WorldSearchResponse = WorldResponse[];

export interface CampaignRequest {
  name: string;
  universe: string;
  world_theme: string;
  tone: string;
  custom_campaign?: string;
  linked_world_id?: string;
  linked_character_ids?: string[];
}

export interface CampaignResponse {
  id: string;
  name: string;
  universe: string;
  world_theme: string;
  tone: string;
  custom_campaign?: string;
  linked_world_id?: string;
  linked_character_ids?: string[];
  image_filename: string;
}


export type User = {
    id: string;
    name: string;
    email: string;
}

export type UserToken = {
    id: string;
    name: string;
    email: string;
    iat: number;
    exp: number;
}