package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type GetTeamDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep GetTeamDependencies) Do(teamId *int) (*model.Team, error) {

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

func NewGetTeamAdapter(gormAdapter *database.GormAdapter) usecases.GetTeamAdapter {
	return &GetTeamDependencies{gormAdapter}
}
