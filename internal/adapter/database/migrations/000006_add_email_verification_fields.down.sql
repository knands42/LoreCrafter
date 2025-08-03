SELECT cron.unschedule_job('delete_expired_email_verifications_job');

DROP FUNCTION IF EXISTS "delete_expired_email_verifications()";

DROP TRIGGER IF EXISTS "update_users_email_verification_updated_at" ON "users_email_verification";

DROP INDEX IF EXISTS "idx_users_email_verification_token";
DROP INDEX IF EXISTS "idx_users_email_verification_user_id";

DROP TABLE IF EXISTS "users_email_verification";