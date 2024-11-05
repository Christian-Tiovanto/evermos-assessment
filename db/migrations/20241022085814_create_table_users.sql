-- migrate:up
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL DEFAULT '' UNIQUE,
    phone VARCHAR(15) NOT NULL DEFAULT '' UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_on_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_on_phone ON users(phone);

-- migrate:down
DROP TABLE IF EXISTS users;
