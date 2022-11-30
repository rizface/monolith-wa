CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    name varchar NOT NULL,
    username varchar NOT NULL UNIQUE,
    password varchar NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)