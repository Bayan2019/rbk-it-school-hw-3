-- +goose Up
INSERT INTO users (email, password_hash, first_name, last_name, is_active)
VALUES
    ('admin@example.com', '$2a$10$adminhash', 'Admin', 'User', TRUE),
    ('ivan@example.com', '$2a$10$ivanhash', 'Ivan', 'Ivanov', TRUE),
    ('dana@example.com', '$2a$10$danahash', 'Dana', 'Sadykova', TRUE),
    ('ayan@example.com', '$2a$10$ayanhash', 'Ayan', 'Kaliyev', FALSE)
ON CONFLICT DO NOTHING;

-- +goose Down
DELETE FROM users 
WHERE email IN ('admin@example.com', 'ivan@example.com', 'dana@example.com', 'ayan@example.com');