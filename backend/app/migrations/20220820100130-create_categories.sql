-- +migrate Up
CREATE TABLE IF NOT EXISTS categories (
    category_id INTEGER NOT NULL AUTO_INCREMENT,
    category_name VARCHAR(127) NOT NULL,
    category_icon VARCHAR(127) DEFAULT NULL,
    created_at DATE NOT NULL,
    updated_at DATE DEFAULT NULL,
    deleted_at DATE DEFAULT NULL,
    PRIMARY KEY (category_id)
);

-- +migrate Down
DROP TABLE IF EXISTS categories;
