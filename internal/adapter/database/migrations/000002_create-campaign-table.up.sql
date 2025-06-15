CREATE TABLE campaigns (
    id UUID PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    setting_summary TEXT,
    setting TEXT,
    image_url VARCHAR(255),
    is_public BOOLEAN NOT NULL DEFAULT false,
    invite_code VARCHAR(12) UNIQUE,
    created_by UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_campaigns_created_by ON campaigns(created_by);
CREATE INDEX idx_campaigns_invite_code ON campaigns(invite_code);