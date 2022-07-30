-- name: CreateAccount :one
INSERT INTO accounts (
  user_name,
  balance
) VALUES (
  $1, $2
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

-- name: UpdateAccount :one
UPDATE accounts
set balance = $2
where id = $1
RETURNING *;
