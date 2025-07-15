CREATE TABLE "users" (
                       id UUID PRIMARY KEY,
                       username VARCHAR(80) NOT NULL,
                       email VARCHAR(120) NOT NULL,
                       hashed_password VARCHAR(255) NOT NULL,
                       is_active BOOLEAN NOT NULL DEFAULT false,
                       avatar_url TEXT,
                       last_login_at TIMESTAMPTZ,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX IF NOT EXISTS "idx_users_username" ON "users"("username");
CREATE INDEX IF NOT EXISTS "idx_users_email" ON "users"("email");

-- Add a function to update the updated_at column automatically
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
