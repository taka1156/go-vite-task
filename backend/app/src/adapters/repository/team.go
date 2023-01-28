package repository

import (
	"app/entity/model"
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
)

type TeamRepositoryDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func NewTeamRepository(
	sqlAdapter *database.SqlAdapter,
) usecases.TeamRepository {
	return &TeamRepositoryDependencies{sqlAdapter}
}

func (dep TeamRepositoryDependencies) GetTeam(teamId int) (*model.Team, error) {

	// stmt := `
	// select
	// 	ts.team_id
	// 	ts.team_name
	// 	ts.team_icon
	// 	ts.team_memo
	// from
	// 	teams as ts
	// on
	// 	ts.team_id = jut.team_id
	// inner join
	// 	roles as rs
	// on
	// 	rs.role_id = us.role_id
	// where
	// 	ts.team_id = ? and

	// 	teams.deleted_at is null
	// `

	return &model.Team{}, nil
}

func (dep TeamRepositoryDependencies) GetTeams() ([]*model.Team, error) {
	var teams []*model.Team
	//currentTime := time.Now()

	return teams, nil
}

func (dep TeamRepositoryDependencies) CreateTeam(input model.InputTeam) (*int, error) {
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

	team := xo.Team{}

	return &team.TeamID, nil
}

func (dep TeamRepositoryDependencies) UpdateTeam(input model.UpdateTeam) (*int, error) {
	//currentTime := time.Now()

	team := xo.Team{}

	return &team.TeamID, nil
}

func (dep TeamRepositoryDependencies) DeleteTeam(teamId int) error {
	//currentTime := time.Now()

	//updateTeamId := uint(1)

	return nil
}
