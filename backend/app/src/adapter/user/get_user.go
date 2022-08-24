package user

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"database/sql"
)

type GetUserDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep GetUserDependencies) Do(userId *int) (*model.User, error) {

	stmt := `
	select 
		us.user_id,
		us.user_name,
		us.email,
		us.user_icon,
		rs.role_id,
		rs.role_name,
		rs.role_icon
	from
		users as us
	inner join
		roles as rs
	on
		rs.role_id = us.role_id
	where
		users.user_id = ? and

		user.deleted_at is null
	`
	defer dep.sqlAdapter.DB.Close()

	rows, err := dep.sqlAdapter.DB.Query(stmt, userId)

	if err != nil {
		return nil, err
	}

	rows.Next()

	d := struct {
		UserID   string
		UserName string
		Email    string
		UserIcon sql.NullString
		RoleID   string
		RoleName string
		RoleIcon sql.NullString
	}{}

	rows.Scan(
		&d.UserID,
		&d.UserName,
		&d.Email,
		&d.UserIcon,
		&d.RoleID,
		&d.RoleName,
		&d.RoleIcon,
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
			RoleID:   d.RoleID,
			RoleName: d.RoleName,
			RoleIcon: blankToNil(d.RoleIcon),
		},
	}

	return user, nil
}

func NewGetUserAdapter(sqlAdapter *database.SqlAdapter) usecases.GetUserAdapter {
	return &GetUserDependencies{sqlAdapter}
}
