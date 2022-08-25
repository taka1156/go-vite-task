package user

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type UpdateUserDependencies struct {
	gormAdapter *database.GormAdapter
}

func canSelectRole(userId string) {
	// roleIdが存在するか確認
}

func (dep UpdateUserDependencies) Do(input model.UpdateUser) (*int, error) {
	// currentTime := time.Now()

	// stmt := `
	// update
	// 	users
	// set
	// 	user_name = ?
	// 	email = ?
	// 	password = ?
	// 	user_icon = ?
	// 	role_id = ?
	// 	update_at = ?
	// where
	// 	users.user_id = ? and

	// 	users.deleted_at is null
	// `

	id := 1

	return &id, nil
}

func NewUpdateUserAdapter(gormAdapter *database.GormAdapter) usecases.UpdateUserAdapter {
	return &UpdateUserDependencies{gormAdapter}
}
