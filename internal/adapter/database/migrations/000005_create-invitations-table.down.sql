SELECT cron.unschedule('delete_old_expired_campaign_invitations_job');
SELECT cron.unschedule('expire_campaign_invitations_job');

DROP FUNCTION IF EXISTS delete_old_expired_campaign_invitations();
DROP FUNCTION IF EXISTS expire_campaign_invitations();
DROP EXTENSION IF EXISTS pg_cron;

DROP TRIGGER IF EXISTS "update_invitations_updated_at" ON "campaign_invitations";

DROP INDEX IF EXISTS "idx_campaign_invitations_token";
DROP INDEX IF EXISTS "idx_campaign_invitations_user_id";
DROP INDEX IF EXISTS "idx_campaign_invitations_campaign_id";
DROP TABLE IF EXISTS "campaign_invitations";
DROP TYPE IF EXISTS invitation_status;
