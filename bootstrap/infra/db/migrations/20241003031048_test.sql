-- +goose Up
CREATE TABLE IF NOT EXISTS test (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS test;