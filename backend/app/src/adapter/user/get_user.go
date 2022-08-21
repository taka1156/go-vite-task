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
		user.user_id,
		user.user_name,
		user.email
		user.user_icon,
		role_id
	from
		users
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
		RoleID   sql.NullString
	}{}

	rows.Scan(
		&d.UserID,
		&d.UserName,
		&d.Email,
		&d.UserIcon,
		&d.RoleID,
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
		RoleID:   blankToNil(d.RoleID),
	}

	return user, nil
}

func NewGetUserAdapter(sqlAdapter *database.SqlAdapter) usecases.GetUserAdapter {
	return &GetUserDependencies{sqlAdapter}
}
