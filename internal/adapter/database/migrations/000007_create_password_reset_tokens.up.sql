-- Create password_reset_tokens table
CREATE TABLE "password_reset_tokens"
(
    id         UUID PRIMARY KEY,
    user_id    UUID        NOT NULL REFERENCES "users" ("id") ON DELETE CASCADE,
    token      VARCHAR(64) NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    used       BOOLEAN     NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Add indexes for better query performance
CREATE INDEX "idx_password_reset_tokens_user_id" ON "password_reset_tokens" ("user_id");
CREATE INDEX "idx_password_reset_tokens_token" ON "password_reset_tokens" ("token");
CREATE INDEX "idx_password_reset_tokens_expires_at" ON "password_reset_tokens" ("expires_at");

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_password_reset_tokens_updated_at
    BEFORE UPDATE
    ON "password_reset_tokens"
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Delete expired password tokens
CREATE OR REPLACE FUNCTION delete_expired_password_reset_tokens() RETURNS void AS
$$
BEGIN
    DELETE
    FROM password_reset_tokens
    WHERE expires_at < NOW();
END;
$$ LANGUAGE plpgsql;

SELECT cron.schedule(
               'delete_expired_password_reset_tokens_job',
               '0 3 * * *', -- Every day at 03:00 AM
               $$SELECT delete_expired_password_reset_tokens();$$
       );
