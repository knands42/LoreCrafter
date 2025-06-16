// types.ts

export type UserToken = {
  id: string;
  username: string;
  email: string;
  avatar_url?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  exp: number;
};

export type User = {
  id: string;
  username: string;
  email: string;
  avatar_url?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
  last_login_at?: string;
};

export type AuthResponse = {
  token: string;
  expiresAt: string;
  user: User;
};

export type RegisterRequest = {
  username: string;
  email: string;
  password: string;
};

export type LoginRequest = {
  username: string;
  password: string;
};

export type CampaignCreationRequest = {
  title: string;
  setting_summary?: string;
  setting?: string;
  image_url?: string;
  is_public: boolean;
};

export type Campaign = {
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
};

export type CampaignMember = {
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
};
