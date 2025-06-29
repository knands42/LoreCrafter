-- name: CreatePasswordResetToken :one
INSERT INTO password_reset_tokens (id, user_id, token_hash, expires_at)
SELECT 
    @id::uuid,
    u.id,
    @token_hash::varchar,
    @expires_at::timestamptz
FROM users u
WHERE u.email = @email
RETURNING *;

-- name: GetValidPasswordResetToken :one
SELECT * FROM password_reset_tokens
WHERE token_hash = $1
AND used = false
AND expires_at > NOW()
LIMIT 1;

-- name: InvalidatePasswordResetToken :exec
UPDATE password_reset_tokens
SET used = true, updated_at = NOW()
WHERE token_hash = $1;

-- name: InvalidateAllUserTokens :exec
UPDATE password_reset_tokens AS ptr
SET used = true, updated_at = NOW()
FROM users u
WHERE u.id = ptr.user_id
AND u.email = $1;

