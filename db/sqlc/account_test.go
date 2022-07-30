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


	account, err := testQueries.CreateAccount(context.Background(), arg)
	
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.UserName, account.UserName)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
}
