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
    c.id,
    c.title,
    c.setting_summary,
    c.setting,
    c.game_system,
    c.number_of_players,
    c.status,
    c.image_url,
    c.is_public,
    c.invite_code,
    c.setting_metadata,
    c.setting_ai_metadata,
    c.created_by,
    c.created_at,
    c.updated_at
    FROM campaigns as c
                  LEFT JOIN campaign_members as cm
                            ON c.id = cm.campaign_id AND cm.user_id = $2
WHERE c.id = $1 AND (
    c.is_public = true OR cm.user_id IS NOT NULL
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
    c.id,
    c.title,
    c.setting_summary,
    c.setting,
    c.game_system,
    c.number_of_players,
    c.status,
    c.image_url,
    c.is_public,
    c.invite_code,
    c.setting_metadata,
    c.setting_ai_metadata,
    c.created_by,
    c.created_at,
    c.updated_at;

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
SELECT * FROM campaign_members
WHERE campaign_id = $1 AND user_id = $2
LIMIT 1;

-- name: UpdateCampaignMember :one
UPDATE campaign_members
SET 
    role = $3,
    last_accessed = CURRENT_TIMESTAMP
WHERE campaign_id = $1 AND user_id = $2
RETURNING *;

-- name: DeleteCampaignMember :exec
DELETE FROM campaign_members
WHERE campaign_id = $1 AND user_id = $2;

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