-- name: CreateUser :one
INSERT INTO users (
    id,
    username,
    email,
    hashed_password
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByUsernameOrEmail :one
SELECT * FROM users
WHERE username = $1 OR email = $2
LIMIT 1;

-- name: GetUserByID :one
SELECT * from users
WHERE id = @user_id::uuid
LIMIT 1;

-- name: UpdateUserPasswordFromToken :one
UPDATE users AS u
SET hashed_password = $2
FROM password_reset_tokens AS ptr
WHERE u.id = ptr.user_id
AND ptr.token = $1
AND ptr.expires_at > NOW()
AND ptr.used = false
RETURNING *;