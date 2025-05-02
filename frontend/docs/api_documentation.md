# LoreCrafter API Documentation

## Project Overview

LoreCrafter is a tool that helps users create rich backstories for tabletop role-playing game (TTRPG) characters, worlds, and campaigns using AI. The system generates detailed narratives based on user-provided information and can link characters, worlds, and campaigns together to create a cohesive storytelling experience.

This documentation describes the API endpoints for a frontend application that will interact with the LoreCrafter backend.

## API Endpoints

### Characters

#### Create Character

**Endpoint:** `/api/characters`  
**Method:** `POST`  
**Description:** Creates a new character with AI-generated backstory, personality, and appearance.

**Request Payload:**
```json
{
  "name": "Captain Kirk",
  "gender": "male",
  "race": "Human",
  "personality": "Impulsive and passionate",  // Optional
  "appearance": "Short black hair, with big cheeks and long neck",  // Optional
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "custom_story": "Kirk was born in a small village...",  // Optional
  "linked_world_id": "uuid-of-world"  // Optional
}
```

**Response:**
```json
{
  "id": "uuid-string",
  "name": "Captain Kirk",
  "gender": "male",
  "race": "Human",
  "personality": "Impulsive and passionate leader who values loyalty and courage...",
  "appearance": "Captain Kirk has short black hair with big cheeks and a long neck. His piercing blue eyes reflect his determination...",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "backstory": "Born in the small village of Riverdale, Kirk's early years were marked by tragedy...",
  "custom_story": "Kirk was born in a small village...",
  "linked_world_id": "uuid-of-world",
  "image_filename": "character_image_uuid.png"
}
```

#### Search Characters

**Endpoint:** `/api/characters/search`  
**Method:** `GET`  
**Description:** Searches for characters based on a query string.

**Query Parameters:**
- `query` (string, required): The search term
- `top` (integer, optional, default=2): Number of results to return

**Response:**
```json
[
  {
    "id": "uuid-string",
    "name": "Captain Kirk",
    "gender": "male",
    "race": "Human",
    "personality": "Impulsive and passionate leader who values loyalty and courage...",
    "appearance": "Captain Kirk has short black hair with big cheeks and a long neck. His piercing blue eyes reflect his determination...",
    "universe": "D&D",
    "world_theme": "fantasy",
    "tone": "Epic",
    "backstory": "Born in the small village of Riverdale, Kirk's early years were marked by tragedy...",
    "linked_world_id": "uuid-of-world",
    "image_filename": "character_image_uuid.png"
  },
  {
    // Second character result
  }
]
```

#### Get Character

**Endpoint:** `/api/characters/{character_id}`  
**Method:** `GET`  
**Description:** Retrieves a specific character by ID.

**Path Parameters:**
- `character_id` (string, required): The UUID of the character

**Response:**
```json
{
  "id": "uuid-string",
  "name": "Captain Kirk",
  "gender": "male",
  "race": "Human",
  "personality": "Impulsive and passionate leader who values loyalty and courage...",
  "appearance": "Captain Kirk has short black hair with big cheeks and a long neck. His piercing blue eyes reflect his determination...",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "backstory": "Born in the small village of Riverdale, Kirk's early years were marked by tragedy...",
  "linked_world_id": "uuid-of-world",
  "image_filename": "character_image_uuid.png"
}
```

### Worlds

#### Create World

**Endpoint:** `/api/worlds`  
**Method:** `POST`  
**Description:** Creates a new world with AI-generated history and timeline.

**Request Payload:**
```json
{
  "name": "Eldoria",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "custom_history": "Eldoria was founded after the great cataclysm...",  // Optional
  "custom_timeline": "Year 0: The Founding...",  // Optional
  "custom_backstory": "The world is currently in turmoil..."  // Optional
}
```

**Response:**
```json
{
  "id": "uuid-string",
  "name": "Eldoria",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "history": "Eldoria, a realm of ancient magic and forgotten lore, was founded in the aftermath of the Great Cataclysm...",
  "timeline": "Year 0: The Founding of Eldoria after the Great Cataclysm\nYear 100: The Rise of the First Mage King...",
  "custom_history": "Eldoria was founded after the great cataclysm...",
  "custom_timeline": "Year 0: The Founding...",
  "custom_backstory": "The world is currently in turmoil...",
  "image_filename": "world_image_uuid.png"
}
```

#### Search Worlds

**Endpoint:** `/api/worlds/search`  
**Method:** `GET`  
**Description:** Searches for worlds based on a query string.

**Query Parameters:**
- `query` (string, required): The search term
- `top` (integer, optional, default=2): Number of results to return

**Response:**
```json
[
  {
    "id": "uuid-string",
    "name": "Eldoria",
    "universe": "D&D",
    "world_theme": "fantasy",
    "tone": "Epic",
    "history": "Eldoria, a realm of ancient magic and forgotten lore, was founded in the aftermath of the Great Cataclysm...",
    "timeline": "Year 0: The Founding of Eldoria after the Great Cataclysm\nYear 100: The Rise of the First Mage King...",
    "image_filename": "world_image_uuid.png"
  },
  {
    // Second world result
  }
]
```

#### Get World

**Endpoint:** `/api/worlds/{world_id}`  
**Method:** `GET`  
**Description:** Retrieves a specific world by ID.

**Path Parameters:**
- `world_id` (string, required): The UUID of the world

**Response:**
```json
{
  "id": "uuid-string",
  "name": "Eldoria",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "history": "Eldoria, a realm of ancient magic and forgotten lore, was founded in the aftermath of the Great Cataclysm...",
  "timeline": "Year 0: The Founding of Eldoria after the Great Cataclysm\nYear 100: The Rise of the First Mage King...",
  "image_filename": "world_image_uuid.png"
}
```

### Campaigns

#### Create Campaign

**Endpoint:** `/api/campaigns`  
**Method:** `POST`  
**Description:** Creates a new campaign with AI-generated setting and hidden elements.

**Request Payload:**
```json
{
  "name": "The Lost Mines",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "custom_campaign": "The players start in a tavern...",  // Optional
  "linked_world_id": "uuid-of-world",  // Optional
  "linked_character_ids": ["uuid-of-character-1", "uuid-of-character-2"]  // Optional
}
```

**Response:**
```json
{
  "id": "uuid-string",
  "name": "The Lost Mines",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "campaign": "The Lost Mines is an epic adventure set in the forgotten reaches of Eldoria. The campaign begins as the players gather in the Rusty Tankard tavern...",
  "hidden_elements": "1. The tavern keeper is actually a retired adventurer with knowledge of the mines.\n2. The mines contain an ancient artifact of immense power...",
  "custom_campaign": "The players start in a tavern...",
  "linked_world_id": "uuid-of-world",
  "linked_character_ids": ["uuid-of-character-1", "uuid-of-character-2"]
}
```

#### Get Campaign

**Endpoint:** `/api/campaigns/{campaign_id}`  
**Method:** `GET`  
**Description:** Retrieves a specific campaign by ID.

**Path Parameters:**
- `campaign_id` (string, required): The UUID of the campaign

**Response:**
```json
{
  "id": "uuid-string",
  "name": "The Lost Mines",
  "universe": "D&D",
  "world_theme": "fantasy",
  "tone": "Epic",
  "campaign": "The Lost Mines is an epic adventure set in the forgotten reaches of Eldoria. The campaign begins as the players gather in the Rusty Tankard tavern...",
  "hidden_elements": "1. The tavern keeper is actually a retired adventurer with knowledge of the mines.\n2. The mines contain an ancient artifact of immense power...",
  "linked_world_id": "uuid-of-world",
  "linked_character_ids": ["uuid-of-character-1", "uuid-of-character-2"]
}
```

### Assets

#### Get Character Image

**Endpoint:** `/api/assets/character/{image_filename}`  
**Method:** `GET`  
**Description:** Retrieves a character image by filename.

**Path Parameters:**
- `image_filename` (string, required): The filename of the character image

**Response:**
- Image file (PNG)

#### Get World Image

**Endpoint:** `/api/assets/world/{image_filename}`  
**Method:** `GET`  
**Description:** Retrieves a world image by filename.

**Path Parameters:**
- `image_filename` (string, required): The filename of the world image

**Response:**
- Image file (PNG)

### PDF Generation

#### Generate Character PDF

**Endpoint:** `/api/pdf/character/{character_id}`  
**Method:** `GET`  
**Description:** Generates a PDF for a character.

**Path Parameters:**
- `character_id` (string, required): The UUID of the character

**Response:**
- PDF file

#### Generate World PDF

**Endpoint:** `/api/pdf/world/{world_id}`  
**Method:** `GET`  
**Description:** Generates a PDF for a world.

**Path Parameters:**
- `world_id` (string, required): The UUID of the world

**Response:**
- PDF file

## Data Models

### Character

```json
{
  "id": "string (UUID)",
  "name": "string",
  "gender": "string",
  "race": "string",
  "personality": "string",
  "appearance": "string",
  "universe": "string",
  "world_theme": "string",
  "tone": "string",
  "backstory": "string",
  "custom_story": "string (optional)",
  "linked_world_id": "string (UUID, optional)",
  "image_filename": "string (optional)"
}
```

### World

```json
{
  "id": "string (UUID)",
  "name": "string",
  "universe": "string",
  "world_theme": "string",
  "tone": "string",
  "history": "string",
  "timeline": "string",
  "custom_history": "string (optional)",
  "custom_timeline": "string (optional)",
  "custom_backstory": "string (optional)",
  "image_filename": "string (optional)"
}
```

### Campaign

```json
{
  "id": "string (UUID)",
  "name": "string",
  "universe": "string",
  "world_theme": "string",
  "tone": "string",
  "campaign": "string",
  "hidden_elements": "string",
  "custom_campaign": "string (optional)",
  "linked_world_id": "string (UUID, optional)",
  "linked_character_ids": "array of strings (UUIDs, optional)"
}
```

## Error Handling

All API endpoints return standard HTTP status codes:

- `200 OK`: The request was successful
- `400 Bad Request`: The request was invalid or missing required parameters
- `404 Not Found`: The requested resource was not found
- `500 Internal Server Error`: An error occurred on the server

Error responses include a JSON body with details:

```json
{
  "error": "Error message",
  "details": "Additional details about the error (optional)"
}
```