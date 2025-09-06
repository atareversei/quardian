-- +migrate Up
ALTER TABLE users
ADD COLUMN first_name VARCHAR(255) NULL,
ADD COLUMN last_name VARCHAR(255) NULL,
ADD COLUMN birth_date DATE NULL,
ADD COLUMN employee_id VARCHAR(255) UNIQUE NULL;

-- +migrate Down
ALTER TABLE users
DROP COLUMN first_name,
DROP COLUMN last_name,
DROP COLUMN birth_date,
DROP COLUMN employee_id;