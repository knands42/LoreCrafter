-- name: CreatePasswordResetToken :one
INSERT INTO password_reset_tokens (id, user_id, token, expires_at)
SELECT @id::uuid                as id,
       u.id                     as user_id,
       @token::varchar          as token,
       @expires_at::timestamptz as expires_at
FROM users u
WHERE u.email = @email
RETURNING *;

-- name: GetValidPasswordResetToken :one
SELECT *
FROM password_reset_tokens
WHERE token = $1
  AND used = false
  AND expires_at > NOW()
LIMIT 1;

-- name: InvalidatePasswordResetToken :exec
UPDATE password_reset_tokens
SET used       = true,
    updated_at = NOW()
WHERE token = $1;

-- name: InvalidateAllUserTokens :exec
UPDATE password_reset_tokens AS ptr
SET used       = true,
    updated_at = NOW()
FROM users u
WHERE u.id = ptr.user_id
  AND u.email = $1;

