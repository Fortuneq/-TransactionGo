-- name: CreateAccount :one
INSERT INTO accounts (
  id,
  user_name,
  balance
) VALUES (
  $1, $2, $3
)
RETURNING *;

