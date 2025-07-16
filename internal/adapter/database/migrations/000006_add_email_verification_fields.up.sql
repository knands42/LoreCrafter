-- Add email verification fields to users table
CREATE TABLE "users_email_verification" (
    "id" UUID NOT NULL,
    "user_id" UUID NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "email_verification_token" VARCHAR(64) UNIQUE,
    "email_verification_sent_at" TIMESTAMPTZ NOT NULL,
    "email_verification_expires_at" TIMESTAMPTZ NOT NULL,
    "email_verified_at" TIMESTAMPTZ NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);

-- Add index for faster lookups by verification token
CREATE INDEX IF NOT EXISTS "idx_users_email_verification_token" ON "users_email_verification"("email_verification_token");

-- Add index for faster lookups by user id
CREATE INDEX IF NOT EXISTS "idx_users_email_verification_user_id" ON "users_email_verification"("user_id");

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_users_email_verification_updated_at
BEFORE UPDATE ON "users_email_verification"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
