-- +goose Up

CREATE TABLE account_tags (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,

  owner_id INTEGER NOT NULL,
  CONSTRAINT fk_owner_id FOREIGN KEY (owner_id) REFERENCES users(id)
);

-- +goose Down

DROP TABLE account_tags;