-- +goose Up
CREATE TABLE posts (
  id uuid PRIMARY KEY,
  title text NOT NULL,
  url text NOT NULL UNIQUE,
  description text,
  feed_id uuid NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  published_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE posts;
