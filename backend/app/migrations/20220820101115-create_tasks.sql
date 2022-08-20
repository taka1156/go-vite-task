-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks (
    task_id INTEGER NOT NULL AUTO_INCREMENT,
    task_name VARCHAR(127) NOT NULL,
    link VARCHAR(127) DEFAULT NULL,
    level INTEGER NOT NULL,
    description VARCHAR(4000) DEFAULT NULL,
    category_id INTEGER NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE DEFAULT NULL,
    deleted_at DATE DEFAULT NULL,
    PRIMARY KEY (task_id),
    FOREIGN KEY categories_index(category_id) REFERENCES categories(category_id)
);

-- +migrate Down
DROP TABLE IF EXISTS tasks;
