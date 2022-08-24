-- +migrate Up
CREATE TABLE IF NOT EXISTS teams (
    team_id INTEGER NOT NULL AUTO_INCREMENT,
    team_name VARCHAR(50) NOT NULL,
    team_icon VARCHAR(127) DEFAULT NULL,
    team_memo VARCHAR(500) DEFAULT NULL,
    start_date  DATE NOT NULL,
    end_date  DATE DEFAULT NULL,
    is_end BOOLEAN NOT NULL DEFAULT 0,
    created_at DATE NOT NULL,
    updated_at DATE DEFAULT NULL,
    deleted_at DATE DEFAULT NULL,
    PRIMARY KEY (team_id)
);

-- +migrate Down
DROP TABLE IF EXISTS teams;
