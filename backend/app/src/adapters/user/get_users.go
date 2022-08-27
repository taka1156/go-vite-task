package user

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type GetUsersDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep GetUsersDependencies) Do() ([]*model.User, error) {
	var users []*model.User
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

	return users, nil
}

func NewGetUsersAdapter(gormAdapter *database.GormAdapter) usecases.GetUsersAdapter {
	return &GetUsersDependencies{gormAdapter}
}
