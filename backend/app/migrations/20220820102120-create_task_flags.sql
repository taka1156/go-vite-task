-- +migrate Up
CREATE TABLE IF NOT EXISTS task_flags (
    task_flag_id INTEGER NOT NULL AUTO_INCREMENT,
    task_flag_name VARCHAR(50) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (task_flag_id)
);

-- +migrate Down
DROP TABLE IF EXISTS task_flags;
