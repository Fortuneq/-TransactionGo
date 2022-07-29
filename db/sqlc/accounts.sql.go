// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: accounts.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  id,
  user_name,
  balance
) VALUES (
  $1, $2, $3
)
RETURNING id, user_name, balance
`

type CreateAccountParams struct {
	ID       int32   `json:"id"`
	UserName string  `json:"user_name"`
	Balance  float32 `json:"balance"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.ID, arg.UserName, arg.Balance)
	var i Accounts
	err := row.Scan(&i.ID, &i.UserName, &i.Balance)
	return i, err
}
