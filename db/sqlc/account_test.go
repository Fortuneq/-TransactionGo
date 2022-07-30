package db

import (
	"context"
	"testing"
	"transactions/util"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		UserName: util.RandomUser(),
		Balance:  util.RandomMoney(),
	}
	account, err := TestQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.UserName, account.user_name)
	require.Equal(t, arg.Balance, account.balance)

	require.NotZero(t, account.ID)
}
