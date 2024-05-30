-- name: CreateAccountTag :one
INSERT INTO account_tags (id, name, description, owner_id)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: ListAccountTagsByUser :many
SELECT * FROM account_tags WHERE owner_id = ?;

-- name: GetAccountTagByUserAndID :one
SELECT * FROM account_tags WHERE owner_id = ? AND id = ?;

-- name: ListAccountTagsByAccount :many
SELECT * FROM accounts_account_tags WHERE account_id = ? AND status = 'active';