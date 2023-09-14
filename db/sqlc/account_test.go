package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yigithancolak/monke-bank-api/util"
)

func TestCreateAccount(t *testing.T) {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		ID:           uuid.New(),
		Owner:        user.ID,
		Balance:      util.RandomMoney(),
		CurrencyCode: util.RandomCurrency(),
	}

	account, err := testStore.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.ID, account.ID)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.CurrencyCode, account.CurrencyCode)
	require.NotZero(t, account.CreatedAt)

}
