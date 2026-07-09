-- +goose Up
CREATE TABLE users (id uuid PRIMARY KEY, created_at timestamp NOT NULL, updated_at timestamp NOT NULL, name TEXT NOT NULL, UNIQUE(name));

-- +goose Down
DROP TABLE users;
