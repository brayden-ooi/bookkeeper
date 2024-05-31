-- +goose Up

CREATE TABLE accounts_account_tags (
  account_id INTEGER NOT NULL,
  account_tag_id INTEGER NOT NULL,

  status TEXT DEFAULT 'active',
  CONSTRAINT check_status CHECK (status = 'active' OR status = 'inactive'),

  PRIMARY KEY (account_id, account_tag_id),
  CONSTRAINT fk_account_id FOREIGN KEY (account_id) REFERENCES accounts(id),
  CONSTRAINT fk_account_tag_id FOREIGN KEY (account_tag_id) REFERENCES account_tags(id)
);

-- +goose Down

DROP TABLE accounts_account_tags;