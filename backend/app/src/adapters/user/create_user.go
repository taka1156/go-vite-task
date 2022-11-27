package user

import (
	"app/entity/model"
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
	"context"
	"database/sql"
	"time"
)

type CreateUserDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep CreateUserDependencies) Do(input model.InputUser) (*int, error) {
	currentTime := time.Now()

	roleId := 1
	user := &xo.User{
		UserName:  input.UserName,
		Email:     input.Email,
		UserIcon:  sql.NullString{String: *input.UserIcon},
		Password:  input.Password,
		RoleID:    sql.NullInt64{Int64: int64(roleId)},
		CreatedAt: currentTime,
		UpdatedAt: sql.NullTime{Time: currentTime},
	}

	err := user.Insert(context.TODO(), dep.sqlAdapter.DB)
	if err != nil {
		return nil, err
	}

	createUserId := int(user.UserID)

	return &createUserId, nil
}

func NewCreateUserAdapter(sqlAdapter *database.SqlAdapter) usecases.CreateUserAdapter {
	return &CreateUserDependencies{sqlAdapter}
}
