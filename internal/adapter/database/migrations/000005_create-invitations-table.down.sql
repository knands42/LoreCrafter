DROP TRIGGER IF EXISTS update_invitations_updated_at ON invitations;

DROP INDEX IF EXISTS idx_invitations_email;
DROP INDEX IF EXISTS idx_invitations_campaign_id;
DROP TABLE IF EXISTS invitations;
DROP TYPE IF EXISTS invitation_status;
