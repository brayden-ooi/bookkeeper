-- +goose Up

CREATE TABLE entries (
  __id INTEGER PRIMARY KEY AUTOINCREMENT, -- internal use
  currency TEXT NOT NULL,
  amount INTEGER NOT NULL, 
  type TEXT NOT NULL
  CONSTRAINT check_type CHECK (type = "debit" OR type = "credit"),
  account_id TEXT NOT NULL,
  account_user_id INTEGER NOT NULL,
  transaction_id INTEGER NOT NULL,

  CONSTRAINT fk_account_id_account_user_id FOREIGN KEY (account_id, account_user_id) REFERENCES accounts(id, user_id),
  CONSTRAINT fk_tx_id FOREIGN KEY (transaction_id) REFERENCES transactions(__id)
);

-- +goose Down

DROP TABLE entries;