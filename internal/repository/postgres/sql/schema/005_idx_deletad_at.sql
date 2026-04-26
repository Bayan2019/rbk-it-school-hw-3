-- +goose Up
CREATE INDEX IF NOT EXISTS idx_users_deleted_at
ON users (deleted_at);

-- +goose Down
DROP INDEX IF EXISTS idx_users_deleted_at;