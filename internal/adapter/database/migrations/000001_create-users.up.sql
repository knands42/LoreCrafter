-- Create users table
CREATE TABLE users (
                       id UUID PRIMARY KEY,
                       username VARCHAR(80) NOT NULL,
                       email VARCHAR(120) NOT NULL,
                       hashed_password VARCHAR(255) NOT NULL,
                       is_active BOOLEAN NOT NULL DEFAULT false,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- Create index on username and email for faster lookups
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);