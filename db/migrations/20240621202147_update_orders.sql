-- +goose Up
ALTER TABLE orders ADD COLUMN packaging VARCHAR(16) NOT NULL;
ALTER TABLE orders ADD CONSTRAINT fk_packaging FOREIGN KEY (packaging) REFERENCES packaging(type);

-- +goose Down
ALTER TABLE orders DROP COLUMN packaging;
