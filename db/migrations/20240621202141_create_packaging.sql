-- +goose Up
CREATE TABLE packaging (
    type VARCHAR(16) PRIMARY KEY,
    cost NUMERIC(10,2) NOT NULL,
    weight_limit FLOAT
);

INSERT INTO packaging (type, cost, weight_limit) VALUES
('bag', 5, 10),
('box', 20, 30),
('wrap', 1, NULL);

-- +goose Down
DROP TABLE packaging;
