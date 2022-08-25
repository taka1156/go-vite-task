-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER NOT NULL AUTO_INCREMENT,
    user_icon VARCHAR(127) DEFAULT NULL,
    user_name VARCHAR(50) NOT NULL,
    email VARCHAR(127) NOT NULL UNIQUE,
    password VARCHAR(127) NOT NULL,
    role_id INTEGER DEFAULT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (user_id),
    FOREIGN KEY role_index(role_id) REFERENCES roles(role_id)
);

-- +migrate Down
DROP TABLE IF EXISTS users;
