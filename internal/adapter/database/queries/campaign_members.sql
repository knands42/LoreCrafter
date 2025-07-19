-- name: CreateFirstCampaignMember :one
INSERT INTO campaign_members (
    id,
    campaign_id,
    user_id,
    role
)VALUES (
            $1, $2, $3, 'gm'
        )
RETURNING *;

-- name: CreateCampaignPlayerMember :one
INSERT INTO campaign_members (
    id,
    campaign_id,
    user_id,
    role
)
SELECT
    @id::uuid,
    @campaign_id::uuid,
    @user_id::uuid,
    @role::member_role
FROM campaign_members cm
WHERE cm.user_id = @requester_id::uuid
  AND cm.campaign_id = @campaign_id::uuid
RETURNING *;

-- name: GetCampaignMember :one
SELECT cm.*
FROM campaign_members cm
WHERE cm.campaign_id = @campaign_id::uuid
AND cm.user_id = @member_id::uuid
AND EXISTS (
    SELECT 1 FROM campaign_members cm2
    WHERE cm2.campaign_id = cm.campaign_id
    AND cm2.user_id = @requester_id::uuid
)
LIMIT 1;


-- name: GetCampaignMembers :many
SELECT cm.*
FROM campaign_members cm
WHERE cm.campaign_id = @campaign_id::uuid
  AND EXISTS (
    SELECT 1 FROM campaign_members cm2
    WHERE cm2.campaign_id = cm.campaign_id
      AND cm2.user_id = @requester_id::uuid
)
LIMIT 10;
