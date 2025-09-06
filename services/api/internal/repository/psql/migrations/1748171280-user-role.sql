-- +migrate Up
CREATE TABLE
  user_roles (
    user_id INTEGER REFERENCES users (user_id),
    role_id INTEGER REFERENCES roles (role_id),
    PRIMARY KEY (user_id, role_id),
    created_at TIMESTAMP DEFAULT now (),
    updated_at TIMESTAMP,
    status record_status NOT NULL DEFAULT 'active'
  );

-- +migrate Down
DROP TABLE IF EXISTS user_roles;