DROP TRIGGER IF EXISTS "update_users_email_verification_updated_at" ON "users_email_verification";

DROP INDEX IF EXISTS "idx_users_email_verification_token";
DROP INDEX IF EXISTS "idx_users_email_verification_user_id";

DROP TABLE IF EXISTS "users_email_verification";