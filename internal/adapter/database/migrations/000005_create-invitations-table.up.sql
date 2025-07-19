CREATE TYPE invitation_status AS ENUM ('pending', 'accepted', 'rejected', 'expired');

CREATE TABLE "campaign_invitations" (
    id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL REFERENCES "campaigns"("id") ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    invited_by UUID NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    token VARCHAR(64) NOT NULL UNIQUE,
    status invitation_status NOT NULL DEFAULT 'pending',
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX "idx_campaign_invitations_campaign_id" ON "campaign_invitations"("campaign_id");
CREATE INDEX "idx_campaign_invitations_user_id" ON "campaign_invitations"("user_id");
CREATE INDEX "idx_campaign_invitations_token" ON "campaign_invitations"("token");

-- Create a trigger to update the updated_at column
CREATE TRIGGER update_invitations_updated_at
BEFORE UPDATE ON "campaign_invitations"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
