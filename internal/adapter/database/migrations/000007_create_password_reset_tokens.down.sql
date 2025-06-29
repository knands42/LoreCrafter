DROP TRIGGER IF EXISTS update_password_reset_tokens_updated_at ON password_reset_tokens;

-- Drop indexes
DROP INDEX IF EXISTS idx_password_reset_tokens_expires_at;
DROP INDEX IF EXISTS idx_password_reset_tokens_token_hash;
DROP INDEX IF EXISTS idx_password_reset_tokens_user_id;

-- Finally, drop the table
DROP TABLE IF EXISTS password_reset_tokens;
