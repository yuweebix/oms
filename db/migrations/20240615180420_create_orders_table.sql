-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    stored_until TIMESTAMP NOT NULL,
    return_by TIMESTAMP NOT NULL,
    status INT NOT NULL,
    hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
