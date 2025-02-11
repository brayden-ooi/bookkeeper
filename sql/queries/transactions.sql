-- name: CreateDraft :one
INSERT INTO transactions (year, counter, description, created_at, user_id, date)
VALUES (
  strftime('%Y'),
  (SELECT transaction_counter FROM users WHERE users.id = ? LIMIT 1) + 1,
  "",
  unixepoch(),
  ?,
  unixepoch()
)
RETURNING *;

-- name: UpdateDraft :one
UPDATE transactions SET 
  year = ?,
  description = ?,
  status = "active",
  date = ?
WHERE year = strftime('%Y') AND counter = ? AND user_id = ?
RETURNING *;

-- name: UpdateTransaction :one
UPDATE transactions SET
  year = ?,
  description = ?,
  status = ?,
  date = ?
WHERE year = ? AND counter = ? AND user_id = ?
RETURNING *;

-- name: ListTransactionsByUser :many
SELECT * FROM transactions WHERE user_id = ?;

-- name: GetTransactionByUserAndID :one
SELECT * FROM transactions WHERE user_id = ? AND counter = ? AND year = ?;
