-- Add email verification fields to users table
CREATE TABLE users_email_verification (
    id UUID NOT NULL,
    user_id UUID NOT NULL,
    email_verification_token VARCHAR(255),
    email_verification_sent_at TIMESTAMPTZ NOT NULL,
    email_verification_expires_at TIMESTAMPTZ NOT NULL,
    email_verified_at TIMESTAMPTZ NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Add index for faster lookups by verification token
CREATE INDEX IF NOT EXISTS idx_users_email_verification_token ON users_email_verification(email_verification_token);
