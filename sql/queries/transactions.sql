-- name: CreateTransaction :one
INSERT INTO transactions (year, counter, description, created_at, user_id)
VALUES (
  ?,
  (SELECT transaction_counter FROM users WHERE users.id = ? LIMIT 1) + 1,
  ?,
  ?,
  ?
)
RETURNING *;

-- name: ListTransactionsByUser :many
SELECT * FROM transactions WHERE user_id = ?;

-- name: GetTransactionByUserAndID :one
SELECT * FROM transactions WHERE user_id = ? AND counter = ? AND year = ?;