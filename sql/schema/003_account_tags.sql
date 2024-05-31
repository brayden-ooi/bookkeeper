-- +goose Up

CREATE TABLE account_tags (
  id TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  user_id INTEGER NOT NULL,

  PRIMARY KEY (user_id, id),
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down

DROP TABLE account_tags;