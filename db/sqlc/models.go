// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID           uuid.UUID          `json:"id"`
	Owner        uuid.UUID          `json:"owner"`
	Balance      int32              `json:"balance"`
	CurrencyCode string             `json:"currency_code"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

type Currency struct {
	Code   string      `json:"code"`
	Name   string      `json:"name"`
	Symbol pgtype.Text `json:"symbol"`
}

type Entry struct {
	ID        uuid.UUID          `json:"id"`
	AccountID uuid.UUID          `json:"account_id"`
	Amount    int32              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Transfer struct {
	ID            uuid.UUID          `json:"id"`
	FromAccountID uuid.UUID          `json:"from_account_id"`
	ToAccountID   uuid.UUID          `json:"to_account_id"`
	Amount        int32              `json:"amount"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}
