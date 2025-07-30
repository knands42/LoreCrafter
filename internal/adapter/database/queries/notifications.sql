-- name: CreateNotification :one
INSERT INTO notifications (id,
                           user_id,
                           type,
                           payload)
VALUES ($1, $2, $3, $4)
RETURNING *;