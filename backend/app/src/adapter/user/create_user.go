package user

import (
	"app/adapter/auth"
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"time"
)

type CreateUserDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep CreateUserDependencies) Do(input model.InputUser) (*int, error) {
	currentTime := time.Now()

	stmt := `
		insert into users
			(user_name, email, user_icon, password, created_at, update_at)
		values
			(?, ?, ?, ?, ?, ?, ?)
	`
	defer dep.sqlAdapter.DB.Close()

	hashPass := auth.HashMd5(input.Password)
	result, err := dep.sqlAdapter.DB.Exec(
		stmt,
		input.UserName,
		input.Email,
		input.UserIcon,
		hashPass,
		currentTime,
		currentTime,
	)
	if err != nil {
		return nil, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	castUserId := int(lastInsertId)

	return &castUserId, nil
}

func NewCreateUserAdapter(sqlAdapter *database.SqlAdapter) usecases.CreateUserAdapter {
	return &CreateUserDependencies{sqlAdapter}
}
