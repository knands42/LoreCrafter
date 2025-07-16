-- name: CreateCampaignInvitation :one
INSERT INTO campaign_invitations (id, campaign_id, user_id, invited_by, token, expires_at)
SELECT
        @id::uuid as id,
        @campaign_id::uuid as campaign_id,
        u.id::uuid as user_id,
        @invited_by::uuid as invited_by,
        @token::varchar as token,
        @expires_at::timestamptz as expires_at
FROM users AS u
WHERE u.username = @username
AND NOT EXISTS (
    SELECT 1
    FROM campaign_members as cm
    WHERE cm.user_id = u.id
    AND cm.campaign_id = @campaign_id
) RETURNING *;


-- name: InvalidateAllCampaignInvitations :exec
UPDATE campaign_invitations AS ci
SET status = 'rejected'
FROM users u
WHERE u.id = ci.user_id
AND ci.user_id = @user_id;

-- name: UpdateCampaignInviteStatus :exec
UPDATE campaign_invitations AS ci
SET status = $1
AND ci.token = $2;