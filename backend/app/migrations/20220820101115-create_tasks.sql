-- +migrate Up
CREATE TABLE IF NOT EXISTS tasks (
    task_id INTEGER NOT NULL AUTO_INCREMENT,
    task_name VARCHAR(127) NOT NULL,
    link VARCHAR(127) DEFAULT NULL,
    level INTEGER NOT NULL DEFAULT 1,
    description VARCHAR(4000) DEFAULT NULL,
    is_review Boolean NOT NULL DEFAULT 0,
    is_all_required Boolean NOT NULL DEFAULT 0,
    category_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (task_id),
    FOREIGN KEY categories_index(category_id) REFERENCES categories(category_id)
);

-- +migrate Down
DROP TABLE IF EXISTS tasks;
