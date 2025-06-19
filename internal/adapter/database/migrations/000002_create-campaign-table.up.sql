CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE TYPE game_system_enum AS ENUM (
    'DND_5E',
    'PATHFINDER_2E',
    'COC_7E',
    'OTHER'
);

CREATE TYPE campaign_status_enum AS ENUM (
    'PLANNING',
    'ACTIVE',
    'PAUSED',
    'FINISHED',
    'ARCHIVED'
    );

CREATE TABLE campaigns (
    id UUID PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    setting_summary TEXT NULL,
    setting TEXT NULL,
    game_system game_system_enum NOT NULL,
    number_of_players SMALLINT DEFAULT 1,
    status campaign_status_enum NOT NULL DEFAULT 'PLANNING',
    image_url VARCHAR(255),
    is_public BOOLEAN NOT NULL DEFAULT false,
    invite_code VARCHAR(12) UNIQUE,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- create index for title using gin index
CREATE INDEX idx_campaigns_title_trgm ON campaigns USING GIN (title gin_trgm_ops);

CREATE INDEX idx_campaigns_created_by ON campaigns(created_by);
CREATE INDEX idx_campaigns_invite_code ON campaigns(invite_code);
CREATE INDEX idx_campaigns_status ON campaigns(status);
CREATE INDEX idx_campaigns_game_system ON campaigns(game_system);
CREATE INDEX idx_campaigns_is_public ON campaigns(is_public);
CREATE INDEX idx_campaigns_status_public_system ON campaigns(status, is_public, game_system);

COMMENT ON TABLE campaigns IS 'Campaigns are the main entity in the system. They are used to store the campaign data.';
COMMENT ON COLUMN campaigns.setting_summary IS 'A summary of the campaign setting.';
COMMENT ON COLUMN campaigns.setting IS 'The detailed setting of the campaign.';
COMMENT ON COLUMN campaigns.game_system IS 'The game system used for the campaign (e.g., dnd, pathfinder, etc.).';
COMMENT ON COLUMN campaigns.status IS 'Lifecycle status of the campaign (e.g., planning, active, paused, etc.).';
COMMENT ON COLUMN campaigns.image_url IS 'The URL of the campaign image.';
COMMENT ON COLUMN campaigns.is_public IS 'Whether the campaign is available to players outside the campaign.';
COMMENT ON COLUMN campaigns.invite_code IS 'The invite code for the campaign.';
