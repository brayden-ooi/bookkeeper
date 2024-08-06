-- +goose Up

ALTER TABLE entries
ADD seq_num INTEGER NOT NULL DEFAULT 1;

-- +goose Down

ALTER TABLE entries
DROP seq_num;
