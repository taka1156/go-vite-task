package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"database/sql"
)

type GetTeamUsersDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep GetTeamUsersDependencies) Do(teamId *int) ([]*model.User, error) {
	var users []*model.User

	stmt := `
	select
		us.user_id,
		us.user_name,
		us.email
		us.user_icon,
		rs.role_id
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

	for rows.Next() {
		d := struct {
			UserID   string
			UserName string
			Email    string
			UserIcon sql.NullString
			RoleId   string
			RoleName string
			RoleIcon sql.NullString
		}{}

		rows.Scan(
			&d.UserID,
			&d.UserName,
			&d.Email,
			&d.UserIcon,
			&d.RoleId,
			&d.RoleName,
			&d.RoleIcon,
		)

		blankToNil := func(column sql.NullString) *string {
			if column.Valid {
				return &column.String
			}

			return nil
		}

		user := &model.User{
			UserID:   d.UserID,
			UserName: d.UserName,
			Email:    d.Email,
			UserIcon: blankToNil(d.UserIcon),
			Role: &model.Role{
				RoleID:   d.RoleId,
				RoleName: d.RoleName,
				RoleIcon: blankToNil(d.RoleIcon),
			},
		}
		users = append(users, user)

	}

	return users, nil
}

func NewGetTeamUsersAdapter(sqlAdapter *database.SqlAdapter) usecases.GetTeamUsersAdapter {
	return &GetTeamUsersDependencies{sqlAdapter}
}
