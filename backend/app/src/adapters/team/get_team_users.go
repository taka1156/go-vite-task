package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type GetTeamUsersDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep GetTeamUsersDependencies) Do(teamId *uint) ([]*model.User, error) {
	var users []*model.User

	// stmt := `
	// select
	// 	us.user_id,
	// 	us.user_name,
	// 	us.email
	// 	us.user_icon,
	// 	rs.role_id
	// 	rs.role_name,
	// 	rs.role_icon,
	// from
	// 	j_users_teams as jut
	// inner join
	// 	teams as ts
	// on
	// 	ts.team_id = jut.team_id
	// inner join
	// 	users as us
	// on
	// 	us.user_id = jut.user_id
	// inner join
	// 	roles as rs
	// on
	// 	rs.role_id = us.role_id
	// where
	// 	ts.team_id = ? and

	// 	user.deleted_at is null
	// 	teams.deleted_at is null
	// `

	return users, nil
}

func NewGetTeamUsersAdapter(sqlAdapter *database.SqlAdapter) usecases.GetTeamUsersAdapter {
	return &GetTeamUsersDependencies{sqlAdapter}
}
