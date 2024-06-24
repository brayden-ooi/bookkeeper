-- +goose Up

CREATE TABLE accounts_account_tags (
  account_id TEXT NOT NULL,
  account_user_id INTEGER NOT NULL,
  account_tag_id INTEGER NOT NULL,

  status TEXT NOT NULL DEFAULT 'active',
  CONSTRAINT check_status CHECK (status = 'active' OR status = 'inactive'),

  PRIMARY KEY (account_id, account_user_id, account_tag_id),
  CONSTRAINT fk_account_id_account_user_id FOREIGN KEY (account_id, account_user_id) REFERENCES accounts(id, user_id),
  CONSTRAINT fk_account_tag_id FOREIGN KEY (account_user_id, account_tag_id) REFERENCES account_tags(user_id, id)
);

-- +goose Down

DROP TABLE accounts_account_tags;