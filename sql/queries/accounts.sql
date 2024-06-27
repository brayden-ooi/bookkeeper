-- name: CreateAccount :one
INSERT INTO accounts (id, name, type, user_id) 
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: AttachAccountTag :one
INSERT INTO accounts_account_tags (account_id, account_user_id, account_tag_id, status)
VALUES (?, ?, ?, "active")
RETURNING *;

-- name: ListAccountsByUser :many
SELECT * FROM accounts WHERE user_id = ?;

-- name: GetAccountByUserAndID :one
SELECT * FROM accounts WHERE user_id = ? AND id = ?;

-- name: DeleteAccountByUserAndID :exec
DELETE FROM accounts WHERE user_id = ? AND id = ?;