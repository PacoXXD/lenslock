CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    -- user_name TEXT,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT
);
