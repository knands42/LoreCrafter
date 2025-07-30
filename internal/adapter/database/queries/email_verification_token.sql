-- name: CreateEmailVerificationToken :one
INSERT INTO users_email_verification (id,
                                      user_id,
                                      email_verification_token,
                                      email_verification_sent_at,
                                      email_verification_expires_at,
                                      created_at,
                                      updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;