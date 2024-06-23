-- name: CreateEntry :one
INSERT INTO entries (currency, amount, type, account_id, account_user_id, transaction_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetEntriesByTx :many
SELECT * FROM entries WHERE transaction_id = ?;
