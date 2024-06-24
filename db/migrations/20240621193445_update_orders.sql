-- +goose Up
-- +goose StatementBegin
ALTER TABLE orders
ADD COLUMN cost NUMERIC(10, 2),
ADD COLUMN weight NUMERIC(10, 2),
ADD COLUMN packaging VARCHAR(16);

UPDATE orders
SET
  cost = 0,
  weight = 0,
  packaging = 'wrap';

ALTER TABLE orders
ALTER COLUMN cost SET NOT NULL,
ALTER COLUMN weight SET NOT NULL,
ALTER COLUMN packaging SET NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE orders
DROP COLUMN cost,
DROP COLUMN weight,
DROP COLUMN packaging;
-- +goose StatementEnd
