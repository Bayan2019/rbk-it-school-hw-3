-- +goose Up
CREATE INDEX IF NOT EXISTS idx_users_created_at
ON users (created_at);

-- +goose Down
DROP INDEX IF EXISTS idx_users_created_at;