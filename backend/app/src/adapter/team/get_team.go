package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"database/sql"
)

type GetTeamDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep GetTeamDependencies) Do(teamId *int) (*model.Team, error) {
	stmt := `
	select
		ts.team_id,
		ts.team_name,
		ts.team_icon,
		us.user_id,
		us.user_name,
		us.email
		us.user_icon,
		rs.role_name,
		rs.role_icon,
	from
		j_users_teams as jut
	inner join
		teams as ts
	on
		ts.team_id = jut.team_id
	inner join
		users as us
	on
		us.user_id = jut.user_id
	inner join
		roles as rs
	on
		rs.role_id = us.role_id
	where
		ts.team_id = ? and

		user.deleted_at is null
		teams.deleted_at is null
	`
	defer dep.sqlAdapter.DB.Close()

	rows, err := dep.sqlAdapter.DB.Query(stmt, teamId)
	if err != nil {
		return nil, err
	}

	var teams model.Team
	for rows.Next() {
		d := struct {
			TeamID   string
			TeamName string
			TeamIcon sql.NullString
			UserID   string
			UserName string
			Email    string
			UserIcon sql.NullString
			RoleName sql.NullString
			RoleIcon sql.NullString
		}{}

		rows.Scan(
			&d.TeamID,
			&d.TeamName,
			&d.TeamIcon,
			&d.UserID,
			&d.UserName,
			&d.Email,
			&d.UserIcon,
			&d.RoleName,
			&d.RoleIcon,
		)

	}

	return getTeamRecord, nil
}

func NewGetTeamAdapter(sqlAdapter *database.SqlAdapter) usecases.GetTeamAdapter {
	return &GetTeamDependencies{sqlAdapter}
}
