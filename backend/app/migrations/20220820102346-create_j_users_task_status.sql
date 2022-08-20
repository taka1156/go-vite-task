-- +migrate Up
CREATE TABLE IF NOT EXISTS j_users_task_status (
    j_user_task_state_id INTEGER NOT NULL AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    task_state_id INTEGER NOT NULL,
    created_at DATE NOT NULL,
    updated_at DATE DEFAULT NULL,
    deleted_at DATE DEFAULT NULL,
    PRIMARY KEY (j_user_task_state_id),
    FOREIGN KEY j_users_index2(user_id) REFERENCES users(user_id),
    FOREIGN KEY j_task_status_index(task_state_id) REFERENCES task_status(task_state_id)
);

-- +migrate Down
DROP TABLE IF EXISTS j_users_task_status;
