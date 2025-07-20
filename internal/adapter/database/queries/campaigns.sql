-- name: CreateCampaign :one
INSERT INTO campaigns (
    id,
    title,
    game_system,
    number_of_players,
    status,
    setting_summary,
    setting,
    image_url,
    setting_metadata,
    setting_ai_metadata,
    is_public,
    created_by
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING *;

-- name: GetCampaignByID :one
SELECT
    c.*
    FROM campaigns as c
                  INNER JOIN campaign_members as cm
                            ON c.id = cm.campaign_id
WHERE c.id = $1 AND (
    c.is_public = true OR cm.user_id = $2
)
LIMIT 1;

-- name: UpdateCampaign :one
UPDATE campaigns as c
SET 
    title = $3,
    setting_summary = $4,
    setting = $5,
    image_url = $6,
    is_public = $7,
    game_system = $8,
    number_of_players = $9,
    status = $10,
    updated_at = CURRENT_TIMESTAMP
FROM campaign_members AS cm
WHERE cm.campaign_id = c.id
  AND c.id = $1
  AND (
      c.is_public = true OR
      cm.user_id = $2
    )
RETURNING
    c.*;

-- name: DeleteCampaign :exec
DELETE FROM campaigns AS c
USING campaign_members AS cm
WHERE cm.campaign_id = c.id
  AND c.id = $1
  AND (
     (cm.user_id = $2 AND cm.role = 'gm'::member_role)
     OR c.created_by = $2
     );

-- name: ListCampaignsByUserID :many
SELECT c.* FROM campaigns c
JOIN campaign_members cm ON c.id = cm.campaign_id
WHERE cm.user_id = $1;

-- name: ListCampaignMembers :many
SELECT * FROM campaign_members
WHERE campaign_id = $1;

-- name: GenerateInviteCode :one
UPDATE campaigns
SET invite_code = $2
WHERE id = $1
RETURNING *;

-- name: GetCampaignByInviteCode :one
SELECT * FROM campaigns
WHERE invite_code = $1
LIMIT 1;