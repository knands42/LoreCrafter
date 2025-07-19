CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'rejected', 'expired');

CREATE TABLE "campaign_invitations" (
    id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL REFERENCES "campaigns"("id") ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    invited_by UUID NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    token VARCHAR(64) NOT NULL UNIQUE,
    status invitation_status NOT NULL DEFAULT 'pending',
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX "idx_campaign_invitations_campaign_id" ON "campaign_invitations"("campaign_id");
CREATE INDEX "idx_campaign_invitations_user_id" ON "campaign_invitations"("user_id");
CREATE INDEX "idx_campaign_invitations_token" ON "campaign_invitations"("token");

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_invitations_updated_at
BEFORE UPDATE ON "campaign_invitations"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

--
-- -- Create a function to set status to expired
-- CREATE OR REPLACE FUNCTION mark_expired_campaign_invitations()
--     RETURNS void AS $$
-- BEGIN
--     UPDATE campaign_invitations
--     SET status = 'expired',
--         updated_at = NOW()
--     WHERE status = 'pending'
--       AND expires_at <= NOW();
-- END;
-- $$ LANGUAGE plpgsql;
--
-- -- Create a function to set delete expired invitations
-- CREATE OR REPLACE FUNCTION delete_expired_campaign_invitations()
--     RETURNS void AS $$
-- BEGIN
--     DELETE FROM campaign_invitations
--     WHERE status = 'expired';
-- END;
-- $$ LANGUAGE plpgsql;
--
-- -- Load the extension
-- CREATE EXTENSION IF NOT EXISTS pg_cron;
--
-- -- Schedule the job (runs every day at 00:05)
-- SELECT cron.schedule(
--                'daily_mark_expired_invitations',
--                '5 0 * * *',  -- minute hour day month day-of-week
--                $$SELECT mark_expired_invitations();$$
--        );
--
-- -- Schedule the function to run every Sunday at 2:00 AM
-- SELECT cron.schedule(
--                'weekly_delete_expired_invitations',
--                '0 2 * * 0',  -- minute hour day month day-of-week (0 = Sunday)
--                $$SELECT delete_expired_campaign_invitations();$$
--        );
