DROP TRIGGER IF EXISTS "update_invitations_updated_at" ON "campaign_invitations";

DROP INDEX IF EXISTS "idx_campaign_invitations_token";
DROP INDEX IF EXISTS "idx_campaign_invitations_user_id";
DROP INDEX IF EXISTS "idx_campaign_invitations_campaign_id";
DROP TABLE IF EXISTS "campaign_invitations";
DROP TYPE IF EXISTS invitation_status;
