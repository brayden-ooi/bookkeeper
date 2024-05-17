-- +goose Up

CREATE TABLE accounts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  balance INTEGER NOT NULL DEFAULT 0,
  type TEXT NOT NULL
  CONSTRAINT check_type CHECK (type = "debit" OR type = "credit"),

  owner_id INTEGER NOT NULL,
  CONSTRAINT fk_owner_id FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE RESTRICT
);

-- +goose Down

DROP TABLE accounts;