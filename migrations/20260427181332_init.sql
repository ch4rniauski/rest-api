-- +goose Up
CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(200) NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    is_completed BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE IF EXISTS tasks;
