package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type TransferTxParams struct {
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        int32     `json:"amount"`
}

type TransferTxResults struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {
	var result TransferTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			ID:            uuid.New(),
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			ID:        uuid.New(),
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			ID:        uuid.New(),
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromAccount, result.ToAccount, err = changeAccountsBalances(ctx, q, arg.FromAccountID, arg.ToAccountID, arg.Amount)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func changeAccountsBalances(ctx context.Context, q *Queries, fromAccountID uuid.UUID, toAccountID uuid.UUID, amount int32) (fromAccount Account, toAccount Account, err error) {
	// Get the fromAccount with locking for update
	fromAccount, err = q.GetAccountForUpdate(ctx, fromAccountID)
	if err != nil {
		return
	}

	// Check if the fromAccount has sufficient balance to transfer
	if fromAccount.Balance < amount {
		err = fmt.Errorf("insufficient balance in account %s", fromAccount.ID)
		return
	}

	// Get the toAccount with locking for update
	toAccount, err = q.GetAccountForUpdate(ctx, toAccountID)
	if err != nil {
		return
	}

	// Adjust the balances
	fromAccount.Balance -= amount
	toAccount.Balance += amount

	// Update the fromAccount's balance in the database
	_, err = q.UpdateAccount(ctx, UpdateAccountParams{
		ID:      fromAccount.ID,
		Balance: fromAccount.Balance,
	})
	if err != nil {
		return
	}

	// Update the toAccount's balance in the database
	_, err = q.UpdateAccount(ctx, UpdateAccountParams{
		ID:      toAccount.ID,
		Balance: toAccount.Balance,
	})
	if err != nil {
		return
	}

	return
}
