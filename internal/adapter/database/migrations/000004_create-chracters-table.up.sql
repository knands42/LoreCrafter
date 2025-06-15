CREATE TABLE characters (
    id UUID PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    race VARCHAR(50),
    class VARCHAR(50),
    level INTEGER NOT NULL DEFAULT 1,
    appearance TEXT,
    personality TEXT,
    backstory TEXT,
    image_url TEXT,
    campaign_id UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    is_npc BOOLEAN NOT NULL DEFAULT false,
    metadata JSONB DEFAULT '{}'::JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_characters_campaign_id ON characters(campaign_id);
CREATE INDEX idx_characters_user_id ON characters(user_id);