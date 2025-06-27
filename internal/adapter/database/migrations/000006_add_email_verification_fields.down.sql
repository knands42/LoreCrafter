-- Remove email verification fields and indexes
DROP INDEX IF EXISTS idx_users_email_verification_token;

DROP TABLE users_email_verification;