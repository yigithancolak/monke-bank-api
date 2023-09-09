-- name: CreateAccount :one
INSERT INTO accounts (id, owner, balance, currency_code, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListAccounts :many
SELECT * FROM accounts
LIMIT $1
OFFSET $2;