-- name: CreateEntry :one
INSERT INTO entries (currency, amount, type, account_id, account_user_id, transaction_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetEntriesByTx :many
SELECT
  l.*,
  r.Name
FROM entries l
LEFT JOIN accounts r ON l.account_id = r.id AND l.account_user_id = r.user_id
WHERE transaction_id = ?;
