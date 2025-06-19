DROP INDEX IF EXISTS idx_campaigns_title_trgm;
DROP INDEX IF EXISTS idx_campaigns_created_by;
DROP INDEX IF EXISTS idx_campaigns_invite_code;
DROP INDEX IF EXISTS idx_campaigns_status;
DROP INDEX IF EXISTS idx_campaigns_game_system;
DROP INDEX IF EXISTS idx_campaigns_is_public;
DROP INDEX IF EXISTS idx_campaigns_status_public_system;

DROP TYPE IF EXISTS campaign_status_enum;
DROP TYPE IF EXISTS game_system_enum;

DROP EXTENSION IF EXISTS pg_trgm;
DROP TABLE IF EXISTS campaigns;