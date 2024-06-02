-- +goose Up

CREATE TABLE transactions (
  __id INTEGER PRIMARY KEY AUTOINCREMENT, -- internal use
  year INTEGER NOT NULL, -- financial year
  counter INTEGER NOT NULL,
  description TEXT NOT NULL,
  created_at INTEGER NOT NULL,
  status TEXT DEFAULT 'draft'
  CONSTRAINT check_status CHECK (status = 'active' OR status = 'inactive' OR status = 'draft'),

  user_id INTEGER NOT NULL,

  CONSTRAINT unique_transaction_id UNIQUE (year, counter, user_id),
  CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT
);

-- +goose Down

DROP TABLE transactions;

