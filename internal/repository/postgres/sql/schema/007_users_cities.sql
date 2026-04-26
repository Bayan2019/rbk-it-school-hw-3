-- +goose Up
CREATE TABLE users_cities(
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    city VARCHAR(255) NOT NULL REFERENCES cities(city) ON DELETE CASCADE,
    added_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, city)
);

-- +goose Down
DROP TABLE clients_psychologists;