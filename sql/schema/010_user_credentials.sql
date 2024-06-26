-- +goose Up

CREATE TABLE user_credentials (
  user_id INTEGER NOT NULL,
  email TEXT UNIQUE NOT NULL,
  password TEXT NOT NULL,
  is_verified INTEGER NOT NULL,
  CONSTRAINT check_verified CHECK (is_verified = 0 OR is_verified = 1),

  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down

DROP TABLE user_credentials;