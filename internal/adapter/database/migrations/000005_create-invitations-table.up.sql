CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'rejected', 'expired');

CREATE TABLE "campaign_invitations"
(
    id          UUID PRIMARY KEY,
    campaign_id UUID              NOT NULL REFERENCES "campaigns" ("id") ON DELETE CASCADE,
    user_id     UUID              NOT NULL REFERENCES "users" ("id") ON DELETE CASCADE,
    invited_by  UUID              NOT NULL REFERENCES "users" ("id") ON DELETE CASCADE,
    token       VARCHAR(64)       NOT NULL UNIQUE,
    status      invitation_status NOT NULL DEFAULT 'pending',
    expires_at  TIMESTAMPTZ       NOT NULL,
    created_at  TIMESTAMPTZ       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS "idx_campaign_invitations_campaign_id" ON "campaign_invitations" ("campaign_id");
CREATE INDEX IF NOT EXISTS "idx_campaign_invitations_user_id" ON "campaign_invitations" ("user_id");
CREATE INDEX IF NOT EXISTS "idx_campaign_invitations_token" ON "campaign_invitations" ("token");
CREATE INDEX IF NOT EXISTS "idx_campaign_invitations_status_expires_at"
    ON "campaign_invitations" ("status", "expires_at");
CREATE INDEX IF NOT EXISTS "idx_campaign_invitations_status" ON "campaign_invitations" ("status");
CREATE INDEX IF NOT EXISTS "idx_campaign_invitations_expires_at" ON "campaign_invitations" ("expires_at");

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_invitations_updated_at
    BEFORE UPDATE
    ON "campaign_invitations"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Add pg cron
CREATE EXTENSION pg_cron;

-- Check and update expired campaign invitations
CREATE OR REPLACE FUNCTION expire_campaign_invitations() RETURNS void AS
$$
BEGIN
    UPDATE campaign_invitations
    SET status     = 'expired',
        updated_at = NOW()
    WHERE status = 'pending'
      AND expires_at < NOW();
END;
$$ LANGUAGE plpgsql;

SELECT cron.schedule(
               'expire_campaign_invitations_job',
               '0 0,12 * * *',
               $$SELECT expire_campaign_invitations();$$
       );

-- Delete old expired campaign invitations
CREATE OR REPLACE FUNCTION delete_old_expired_campaign_invitations() RETURNS void AS
$$
BEGIN
    DELETE
    FROM campaign_invitations
    WHERE expires_at < NOW() - INTERVAL '1 day';
END;
$$ LANGUAGE plpgsql;



SELECT cron.schedule(
               'delete_old_expired_campaign_invitations_job',
               '0 1 * * *', -- Every day at 01:00 AM
               $$SELECT delete_old_expired_campaign_invitations();$$
       );