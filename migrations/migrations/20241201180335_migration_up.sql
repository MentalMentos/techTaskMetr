-- +goose Up
CREATE TABLE "tasks" (
    id VARCHAR(50) NOT NULL PRIMARY KEY UNIQUE,
    title VARCHAR(50) NOT NULL,
    description VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose Down
DROP TABLE  "tasks";
