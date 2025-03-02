-- +goose Up
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    age        INT,
    email      TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE users;


