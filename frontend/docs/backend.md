# LoreCrafter - API Documentation

## Project Overview
LoreCrafter is a collaborative storytelling and campaign management platform designed for tabletop role-playing games (TTRPGs). It provides tools for Game Masters (GMs) and players to create, manage, and participate in immersive campaigns with rich lore and character development.

## Key Features

1. **User Authentication**
   - User registration and login
   - JWT-based authentication
   - User profile management

2. **Campaign Management**
   - Create and manage TTRPG campaigns
   - Public and private campaign settings
   - Rich text support for campaign settings and lore
   - Campaign image uploads

3. **Collaboration**
   - Invite system with unique invite codes
   - Role-based access control (GM and Player roles)
   - Campaign member management

4. **Lore Management**
   - Timeline events tracking
   - Character profiles
   - World-building tools

## API Endpoints

### Authentication

#### `POST /api/auth/register`
Register a new user.

**Request Body:**
```json
{
  "username": "string",
  "email": "string",
  "password": "string"
}
```

**Response (200 OK):**
```json
{
  "token": "string",
  "expiresAt": "string",
  "user": {
    "id": "string",
    "username": "string",
    "email": "string",
    "avatar_url": "string",
    "is_active": true,
    "created_at": "string",
    "updated_at": "string"
  }
}
```

#### `POST /api/auth/login`
Authenticate a user.

**Request Body:**
```json
{
  "username": "string",
  "password": "string"
}
```

**Response (200 OK):**
```json
{
  "token": "string",
  "expiresAt": "string",
  "user": {
    "id": "string",
    "username": "string",
    "email": "string",
    "avatar_url": "string",
    "is_active": true,
    "created_at": "string",
    "updated_at": "string"
  }
}
```

### Campaigns

#### `GET /api/campaigns`
List all campaigns the user has access to.

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
[
  {
    "id": "string",
    "title": "string",
    "setting_summary": "string",
    "image_url": "string",
    "is_public": true,
    "created_at": "string",
    "updated_at": "string",
    "created_by": "string"
  }
]
```

#### `POST /api/campaigns`
Create a new campaign.

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "title": "string",
  "setting_summary": "string",
  "setting": "string",
  "image_url": "string",
  "is_public": true
}
```

**Response (201 Created):**
```json
{
  "id": "string",
  "title": "string",
  "setting_summary": "string",
  "setting": "string",
  "image_url": "string",
  "is_public": true,
  "invite_code": "string",
  "created_by": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

#### `GET /api/campaigns/{campaignID}`
Get campaign details by ID.

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "id": "string",
  "title": "string",
  "setting_summary": "string",
  "setting": "string",
  "image_url": "string",
  "is_public": true,
  "invite_code": "string",
  "created_by": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

### Campaign Members

#### `GET /api/campaigns/{campaignID}/members`
List all members of a campaign.

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
[
  {
    "id": "string",
    "user_id": "string",
    "campaign_id": "string",
    "role": "gm" | "player",
    "joined_at": "string",
    "last_accessed": "string",
    "user": {
      "id": "string",
      "username": "string",
      "email": "string",
      "avatar_url": "string"
    }
  }
]
```

#### `POST /api/campaigns/join/{inviteCode}`
Join a campaign using an invite code.

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "id": "string",
  "campaign_id": "string",
  "user_id": "string",
  "role": "player",
  "joined_at": "string"
}
```

### User Profile

#### `GET /api/me`
Get the current user's profile.

**Headers:**
```
Authorization: Bearer <token>
```

**Response (200 OK):**
```json
{
  "id": "string",
  "username": "string",
  "email": "string",
  "avatar_url": "string",
  "is_active": true,
  "created_at": "string",
  "updated_at": "string"
}
```

## Authentication
All endpoints except `/api/auth/register` and `/api/auth/login` require authentication. Include the JWT token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

## Error Responses

### 400 Bad Request
```json
{
  "error": "string",
  "details": {}
}
```

### 401 Unauthorized
```json
{
  "error": "invalid or expired token"
}
```

### 403 Forbidden
```json
{
  "error": "insufficient permissions"
}
```

### 404 Not Found
```json
{
  "error": "resource not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "internal server error"
}
```

## Data Types

### User
```typescript
interface User {
  id: string;
  username: string;
  email: string;
  avatar_url?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  last_login_at?: string;
}
```

### Campaign
```typescript
interface Campaign {
  id: string;
  title: string;
  setting_summary?: string;
  setting?: string;
  image_url?: string;
  is_public: boolean;
  invite_code?: string;
  created_by: string;
  created_at: string;
  updated_at: string;
}
```

### Campaign Member
```typescript
interface CampaignMember {
  id: string;
  campaign_id: string;
  user_id: string;
  role: 'gm' | 'player';
  joined_at: string;
  last_accessed?: string;
  user?: {
    id: string;
    username: string;
    email: string;
    avatar_url?: string;
  };
}
```

## Frontend Implementation Notes

1. **Authentication Flow**
   - Store the JWT token in memory or secure storage (e.g., HTTP-only cookies)
   - Implement token refresh logic
   - Redirect unauthenticated users to login

2. **Error Handling**
   - Handle 401 errors by redirecting to login
   - Display user-friendly error messages
   - Implement retry logic for failed requests

3. **Real-time Updates**
   - Consider implementing WebSocket for real-time updates
   - Poll for updates on campaign changes if needed

4. **File Uploads**
   - Use presigned URLs for secure file uploads
   - Validate file types and sizes client-side

5. **Responsive Design**
   - Ensure the UI works on both desktop and mobile devices
   - Use a responsive grid system for campaign cards and other components

## Development Setup

1. Clone the repository
2. Install dependencies
3. Set up environment variables (see `.env.example`)
4. Run migrations
5. Start the development server

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a pull request

## License
Apache 2.0

---

This documentation is auto-generated based on the API specification. For the most up-to-date information, refer to the actual API responses and Swagger documentation.
