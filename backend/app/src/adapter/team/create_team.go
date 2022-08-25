package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type CreateTeamDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep CreateTeamDependencies) Do(input model.InputTeam) (*int, error) {
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

	id := 1

	return &id, nil
}

func NewCreateTeamAdapter(gormAdapter *database.GormAdapter) usecases.CreateTeamAdapter {
	return &CreateTeamDependencies{gormAdapter}
}
