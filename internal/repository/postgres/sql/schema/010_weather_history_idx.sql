-- +goose Up
CREATE INDEX IF NOT EXISTS idx_weather_history_user_city 
ON users_cities (user_id, city);

-- +goose Down
DROP INDEX IF EXISTS idx_weather_history_user_city;