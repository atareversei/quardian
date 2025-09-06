-- +migrate Up
CREATE TABLE
    users (
        user_id SERIAL PRIMARY KEY,
        username VARCHAR(32) NOT NULL UNIQUE,
        email VARCHAR(100) UNIQUE,
        mobile VARCHAR(15) UNIQUE,
        password_hash VARCHAR(64) NOT NULL,
        created_at TIMESTAMP DEFAULT now (),
        updated_at TIMESTAMP,
        status record_status NOT NULL DEFAULT 'pending'
    );

-- +migrate Down
DROP TABLE IF EXISTS users;