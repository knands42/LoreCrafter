-- name: CreateCampaignMember :one
INSERT INTO campaign_members (
    id,
    campaign_id,
    user_id,
    role
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

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
