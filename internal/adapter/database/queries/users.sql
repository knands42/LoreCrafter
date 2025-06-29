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

-- name: UpdateUserPassword :one
UPDATE users
SET hashed_password = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserPasswordFromToken :one
UPDATE users AS u
SET hashed_password = $2
FROM password_reset_tokens AS ptr
WHERE u.id = ptr.user_id
AND ptr.token_hash = $1
RETURNING *;