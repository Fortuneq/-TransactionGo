-- name: CreateAccount :one
INSERT INTO accounts (
  id,
  user_name,
  balance
) VALUES (
  $1, $2, $3
)
RETURNING *;


-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY name;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
