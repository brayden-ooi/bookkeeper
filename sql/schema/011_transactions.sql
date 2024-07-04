-- +goose Up

ALTER TABLE transactions
ADD date INTEGER NOT NULL;

-- +goose Down

ALTER TABLE transactions
DROP date;
