-- name: CreateTransfer :one
INSERT INTO transfers (id,from_account_id,to_account_id,amount)
VALUES ($1, $2, $3, $4)
RETURNING *;