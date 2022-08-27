-- +migrate Up
CREATE TABLE IF NOT EXISTS j_users_teams (
    j_user_team_id INTEGER NOT NULL AUTO_INCREMENT,
    user_id INTEGER NOT NULL,
    team_id INTEGER NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME DEFAULT NULL,
    deleted_at DATETIME DEFAULT NULL,
    PRIMARY KEY (j_user_team_id),
    FOREIGN KEY j_users_index2(user_id) REFERENCES users(user_id),
    FOREIGN KEY j_teams(team_id) REFERENCES teams(team_id)
);

-- +migrate Down
DROP TABLE IF EXISTS  j_users_teams;
