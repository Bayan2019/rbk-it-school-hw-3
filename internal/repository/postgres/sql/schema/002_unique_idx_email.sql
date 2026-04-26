-- +goose Up
CREATE UNIQUE INDEX IF NOT EXISTS ux_users_email
ON users (LOWER(email))
WHERE deleted_at IS NULL;

-- +goose Down
DROP INDEX IF EXISTS ux_users_email;