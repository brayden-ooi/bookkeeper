-- +goose Up

CREATE TABLE accounts (
  id TEXT NOT NULL,
  name TEXT NOT NULL,
  balance INTEGER NOT NULL DEFAULT 0,
  type TEXT NOT NULL
  CONSTRAINT check_type CHECK (type = "debit" OR type = "credit"),
  user_id INTEGER NOT NULL,

  PRIMARY KEY (id, user_id),
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);

-- +goose Down

DROP TABLE accounts;