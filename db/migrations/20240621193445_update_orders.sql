-- +goose Up
ALTER TABLE orders ADD COLUMN cost FLOAT NOT NULL;
ALTER TABLE orders ADD COLUMN weight FLOAT NOT NULL;

-- +goose Down
ALTER TABLE orders DROP COLUMN cost;
ALTER TABLE orders DROP COLUMN weight;