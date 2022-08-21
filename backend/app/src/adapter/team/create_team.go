package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"time"
)

type CreateTeamDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep CreateTeamDependencies) Do(input model.InputTeam) (*int, error) {
	currentTime := time.Now()

	stmt_teams := `
	insert into teams
		(team_name, team_icon, team_memo, start_date, end_date created_at, update_at)
	values
		(?, ?, ?, ?, ?)
	`

	stmt_j_users_teams := `
	insert into j_users_teams
		(user_id, teams_id, created_at, update_at)
	values
		(?, ?, ?, ?)
	`

	defer dep.sqlAdapter.DB.Close()

	ts, err := dep.sqlAdapter.DB.Begin()
	if err != nil {
		return nil, err
	}

	result, err := ts.Exec(
		stmt_teams,
		input.TeamName,
		input.TeamIcon,
		input.TeamMemo,
		input.StartDate,
		input.EndDate,
		currentTime,
		currentTime,
	)

	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	castTeamId := int(lastInsertId)

	for _, userId := range input.Users {
		_, err := ts.Exec(
			stmt_j_users_teams,
			castTeamId,
			userId,
			currentTime,
			currentTime,
		)

		if err != nil {
			ts.Rollback()
			return nil, err
		}
	}
	ts.Commit()

	return &castTeamId, nil
}

func NewCreateTeamAdapter(sqlAdapter *database.SqlAdapter) usecases.CreateTeamAdapter {
	return &CreateTeamDependencies{sqlAdapter}
}
