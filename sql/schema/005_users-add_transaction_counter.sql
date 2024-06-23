-- +goose Up

ALTER TABLE users ADD transaction_counter INTEGER NOT NULL DEFAULT 0;

-- +goose Down

ALTER TABLE users DROP COLUMN transaction_counter;
