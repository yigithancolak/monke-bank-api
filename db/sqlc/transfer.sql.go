// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: transfer.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (id,from_account_id,to_account_id,amount)
VALUES ($1, $2, $3, $4)
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	ID            uuid.UUID `json:"id"`
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        int32     `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, createTransfer,
		arg.ID,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}