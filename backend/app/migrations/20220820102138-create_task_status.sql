-- +migrate Up
CREATE TABLE IF NOT EXISTS task_status (
    task_state_id INTEGER NOT NULL AUTO_INCREMENT,
    task_id INTEGER NOT NULL,
    task_flag_id INTEGER NOT NULL,
    progress INTEGER NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE DEFAULT NULL,
    deleted_at DATE DEFAULT NULL,
    PRIMARY KEY (task_state_id),
    FOREIGN KEY tasks_index(task_id) REFERENCES tasks(task_id),
    FOREIGN KEY task_flags_index(task_flag_id) REFERENCES task_flags(task_flag_id)
);

-- +migrate Down
DROP TABLE IF EXISTS task_status;
