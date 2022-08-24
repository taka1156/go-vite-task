package user

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"strconv"
	"time"
)

type UpdateUserDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func canSelectRole(userId string) {
	// roleIdが存在するか確認
}

func (dep UpdateUserDependencies) Do(input model.UpdateUser) (*int, error) {
	currentTime := time.Now()

	stmt := `
	update
		users
	set
		user_name = ?
		email = ?
		password = ?
		user_icon = ?
		role_id = ?
		update_at = ?
	where
		users.user_id = ? and

		users.deleted_at is null
	`
	defer dep.sqlAdapter.DB.Close()

	_, err := dep.sqlAdapter.DB.Exec(
		stmt,
		input.UserName,
		input.Email,
		input.Password,
		input.UserIcon,
		input.RoleID,
		currentTime,
		input.UserID,
	)
	if err != nil {
		return nil, err
	}

	castUserId, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, err
	}

	return &castUserId, nil
}

func NewUpdateUserAdapter(sqlAdapter *database.SqlAdapter) usecases.UpdateUserAdapter {
	return &UpdateUserDependencies{sqlAdapter}
}
