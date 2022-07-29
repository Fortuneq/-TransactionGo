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

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, user_name, balance FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Accounts, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Accounts
	err := row.Scan(&i.ID, &i.UserName, &i.Balance)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, user_name, balance FROM accounts
ORDER BY name
`

func (q *Queries) ListAccounts(ctx context.Context) ([]Accounts, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Accounts
	for rows.Next() {
		var i Accounts
		if err := rows.Scan(&i.ID, &i.UserName, &i.Balance); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
