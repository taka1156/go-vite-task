package user

import (
	"app/infra/database"
	"app/usecases"
)

type DeleteUserDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep DeleteUserDependencies) Do(userId int) (*bool, error) {

	// stmt := `
	// select
	// 	us.user_id,
	// 	us.user_name,
	// 	us.email,
	// 	us.user_icon,
	// 	rs.role_id,
	// 	rs.role_name,
	// 	rs.role_icon
	// from
	// 	users as us
	// inner join
	// 	roles as rs
	// on
	// 	rs.role_id = us.role_id
	// where
	// 	users.user_id = ? and

	// 	user.deleted_at is null
	// `

	isDelete := false

	return &isDelete, nil
}

func NewDeleteUserAdapter(gormAdapter *database.GormAdapter) usecases.DeleteUserAdapter {
	return &DeleteUserDependencies{gormAdapter}
}
