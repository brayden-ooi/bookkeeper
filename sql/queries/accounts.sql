-- name: CreateAccount :one
INSERT INTO accounts (name, type, owner_id) 
VALUES (?, ?, ?)
RETURNING *;