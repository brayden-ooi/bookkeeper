-- +goose Up

CREATE TABLE accounts_account_tags (
  account_id INTEGER,
  account_tag_id INTEGER,

  status TEXT DEFAULT "active",
  CONSTRAINT check_status CHECK (status = "active" OR status = "inactive"),

  CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES accounts(id),
  CONSTRAINT fk_account_tag_id FOREIGN KEY (account_tag_id) REFERENCES account_tags(id),

  CONSTRAINT unique_account_tag UNIQUE(account_id, account_tag_id)
);

-- +goose Down

DROP TABLE accounts_account_tags;