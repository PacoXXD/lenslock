-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    -- user_name TEXT,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
