CREATE TYPE member_role AS ENUM ('gm', 'player');

CREATE TABLE "campaign_members"
(
    id            UUID PRIMARY KEY,
    campaign_id   UUID        NOT NULL REFERENCES "campaigns" ("id") ON DELETE CASCADE,
    user_id       UUID        NOT NULL REFERENCES "users" ("id") ON DELETE CASCADE,
    role          member_role NOT NULL DEFAULT 'player',
    joined_at     TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_accessed TIMESTAMPTZ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_campaign_members_campaign_id" ON "campaign_members" ("campaign_id");
CREATE INDEX IF NOT EXISTS "idx_campaign_members_user_id" ON "campaign_members" ("user_id");
CREATE INDEX IF NOT EXISTS "idx_campaign_id_and_user_id" ON "campaign_members" ("campaign_id", "user_id");
CREATE INDEX IF NOT EXISTS "idx_campaign_members_role" ON "campaign_members" ("role");

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_campaign_members_updated_at
    BEFORE UPDATE
    ON "campaign_members"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
