-- +goose Up
CREATE INDEX IF NOT EXISTS idx_users_is_active
ON users (is_active);

-- +goose Down
DROP INDEX IF EXISTS idx_users_is_active;