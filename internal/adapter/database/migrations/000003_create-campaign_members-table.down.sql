DROP TRIGGER IF EXISTS update_campaign_members_updated_at ON campaign_members;


DROP INDEX IF EXISTS idx_campaign_members_user_id;
DROP INDEX IF EXISTS idx_campaign_members_campaign_id;
DROP TABLE IF EXISTS campaign_members;
DROP TYPE IF EXISTS member_role;