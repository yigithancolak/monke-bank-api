-- name: CreateAccount :one
INSERT INTO accounts (id, owner, balance, currency_code)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListAccounts :many
SELECT * FROM accounts
LIMIT $1
OFFSET $2;