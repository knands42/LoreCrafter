CREATE TYPE member_role AS ENUM ('gm', 'player');

CREATE TABLE campaign_members (
    id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role member_role NOT NULL DEFAULT 'player',
    joined_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_accessed TIMESTAMPTZ,
    UNIQUE(campaign_id, user_id)
);

CREATE INDEX idx_campaign_members_campaign_id ON campaign_members(campaign_id);
CREATE INDEX idx_campaign_members_user_id ON campaign_members(user_id);