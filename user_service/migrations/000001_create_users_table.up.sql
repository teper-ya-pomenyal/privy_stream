CREATE TABLE users (
    uuid UUID PRIMARY KEY,
    user_name TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    birth_date DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
)