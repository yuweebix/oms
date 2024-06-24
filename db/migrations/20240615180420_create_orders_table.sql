-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS
    orders (
        id BIGINT PRIMARY KEY,
        user_id BIGINT NOT NULL,
        stored_until TIMESTAMP NOT NULL,
        return_by TIMESTAMP NOT NULL,
        status VARCHAR(16) NOT NULL,
        hash TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );

CREATE INDEX idx_orders_user_id ON orders (user_id); -- для лук-апа по пользователям
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd