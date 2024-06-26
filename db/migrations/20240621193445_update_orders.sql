-- +goose Up
-- +goose StatementBegin
ALTER TABLE IF NOT EXISTS orders
  ADD COLUMN IF NOT EXISTS cost BIGINT NOT NULL DEFAULT 0, -- в микрорублях
  ADD COLUMN IF NOT EXISTS weight REAL NOT NULL DEFAULT 0,
  ADD COLUMN IF NOT EXISTS packaging VARCHAR(16) DEFAULT 'wrap';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE IF EXISTS orders
  DROP COLUMN IF EXISTS cost,
  DROP COLUMN IF EXISTS weight,
  DROP COLUMN IF EXISTS packaging;
-- +goose StatementEnd
