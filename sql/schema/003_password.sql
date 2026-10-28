-- +goose Up
ALTER TABLE users
ADD COLUMN hash_password TEXT NOT NULL
DEFAULT 'unset';

-- +goose Down
ALTER TABLE users
DROP COLUMN hash_password