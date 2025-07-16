CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'rejected', 'expired');

CREATE TABLE "invitations" (
    id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL REFERENCES "campaigns"("id") ON DELETE CASCADE,
    email VARCHAR(120) NOT NULL,
    invited_by UUID NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    token VARCHAR(32) NOT NULL UNIQUE,
    status invitation_status NOT NULL DEFAULT 'pending',
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX "idx_invitations_campaign_id" ON "invitations"("campaign_id");
CREATE INDEX "idx_invitations_email" ON "invitations"("email");

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_invitations_updated_at
BEFORE UPDATE ON "invitations"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
