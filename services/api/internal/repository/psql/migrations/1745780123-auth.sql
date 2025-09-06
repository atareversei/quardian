-- +migrate Up
CREATE TABLE resources
(
    resource_id SERIAL PRIMARY KEY,
    name        VARCHAR(32)   NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP              DEFAULT now(),
    updated_at  TIMESTAMP,
    status      record_status NOT NULL DEFAULT 'active'
);

CREATE TABLE actions
(
    action_id   SERIAL PRIMARY KEY,
    name        VARCHAR(32)   NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP              DEFAULT now(),
    updated_at  TIMESTAMP,
    status      record_status NOT NULL DEFAULT 'active'
);

CREATE TABLE feasible_actions_on_resources
(
    resource_id INTEGER REFERENCES resources (resource_id),
    action_id   INTEGER REFERENCES actions (action_id),
    PRIMARY KEY (resource_id, action_id),
    created_at  TIMESTAMP              DEFAULT now(),
    updated_at  TIMESTAMP,
    status      record_status NOT NULL DEFAULT 'active'
);

CREATE TABLE roles
(
    role_id     SERIAL PRIMARY KEY,
    name        VARCHAR(32)   NOT NULL UNIQUE,
    description TEXT,
    created_at  TIMESTAMP              DEFAULT now(),
    updated_at  TIMESTAMP,
    status      record_status NOT NULL DEFAULT 'active'
);

CREATE TABLE permissions
(
    role_id     INTEGER REFERENCES roles (role_id),
    resource_id INTEGER REFERENCES resources (resource_id),
    action_id   INTEGER REFERENCES actions (action_id),
    PRIMARY KEY (role_id, resource_id, action_id),
    created_at  TIMESTAMP              DEFAULT now(),
    updated_at  TIMESTAMP,
    status      record_status NOT NULL DEFAULT 'active'

);

-- +migrate Down
DROP TABLE IF EXISTS actions;
DROP TABLE IF EXISTS resources;
DROP TABLE IF EXISTS feasible_actions_on_resources;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS permissions;