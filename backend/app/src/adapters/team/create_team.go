package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type CreateTeamDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep CreateTeamDependencies) Do(input model.InputTeam) (*uint, error) {
	// currentTime := time.Now()

	// stmt_teams := `
	// insert into teams
	// 	(team_name, team_icon, team_memo, start_date, end_date created_at, update_at)
	// values
	// 	(?, ?, ?, ?, ?)
	// `

	// stmt_j_users_teams := `
	// insert into j_users_teams
	// 	(user_id, teams_id, created_at, update_at)
	// values
	// 	(?, ?, ?, ?)
	// `

	id := uint(1)

	return &id, nil
}

func NewCreateTeamAdapter(sqlAdapter *database.SqlAdapter) usecases.CreateTeamAdapter {
	return &CreateTeamDependencies{sqlAdapter}
}
