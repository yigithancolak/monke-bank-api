-- name: CreateAccount :one
INSERT INTO accounts (id, owner, balance, currency_code)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListAccounts :many
SELECT * FROM accounts
LIMIT $1
OFFSET $2;

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = $1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1
FOR NO KEY UPDATE;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;
