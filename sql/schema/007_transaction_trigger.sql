-- +goose Up

-- +goose StatementBegin
CREATE TRIGGER increment_transaction_counter
AFTER INSERT ON transactions
BEGIN 
  UPDATE users SET transaction_counter = transaction_counter + 1 WHERE id = NEW.user_id;
END;
-- +goose StatementEnd

-- +goose Down

DROP TRIGGER increment_transaction_counter;
