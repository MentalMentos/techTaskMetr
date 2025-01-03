-- +goose Up
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL ,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    status boolean NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS tasks;
