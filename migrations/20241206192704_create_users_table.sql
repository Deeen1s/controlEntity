-- +goose Up
-- SQL для применения миграции
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name TEXT NOT NULL,
                       email TEXT UNIQUE NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW()

                   );
-- +goose Down
-- SQL для отмены миграции
DROP TABLE users;