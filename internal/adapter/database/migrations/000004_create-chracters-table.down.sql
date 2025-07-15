DROP TRIGGER IF EXISTS "update_characters_updated_at" ON "characters";

DROP INDEX IF EXISTS "idx_characters_user_id";
DROP INDEX IF EXISTS "idx_characters_campaign_id";
DROP TABLE IF EXISTS "characters";
