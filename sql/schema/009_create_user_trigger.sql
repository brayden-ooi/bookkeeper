-- +goose Up

-- +goose StatementBegin
CREATE TRIGGER create_user_trigger
AFTER INSERT ON users
BEGIN 
  INSERT INTO account_tags (id, name, description, user_id) VALUES
  ("_1000", "Assets", "", NEW.id),
  ("_1001", "Current Assets", "", NEW.id),
  ("_1002", "Non-current Assets", "", NEW.id),
  ("_2000", "Liabilities", "", NEW.id),
  ("_2001", "Current Liabilities", "", NEW.id),
  ("_2002", "Non-current Liabilities", "", NEW.id),
  ("_3000", "Equity", "", NEW.id),
  ("_4000", "Expenses", "", NEW.id),
  ("_5000", "Revenue", "", NEW.id);
END;
-- +goose StatementEnd

-- +goose Down

DROP TRIGGER create_user_trigger;
