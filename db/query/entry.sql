-- name: CreateEntry :one
INSERT INTO entries (id,account_id,amount)
VALUES ($1, $2, $3)
RETURNING *;