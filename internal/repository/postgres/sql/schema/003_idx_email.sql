-- +goose Up
CREATE INDEX IF NOT EXISTS idx_users_email
ON users (LOWER(email));

-- +goose Down
DROP INDEX IF EXISTS idx_users_email;