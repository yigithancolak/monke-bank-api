-- name: CreateUser :one
INSERT INTO users (id, email, password, full_name)
VALUES ($1, $2, $3, $4)
RETURNING *;