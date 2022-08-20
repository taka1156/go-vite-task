-- +migrate Up
CREATE TABLE IF NOT EXISTS roles (
    role_id INTEGER NOT NULL AUTO_INCREMENT,
    role_icon VARCHAR(127) DEFAULT NULL,
    role_name VARCHAR(30) NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE DEFAULT NULL,
    deleted_at DATE DEFAULT NULL,
    PRIMARY KEY (role_id)
);

-- +migrate Down
DROP TABLE IF EXISTS roles;