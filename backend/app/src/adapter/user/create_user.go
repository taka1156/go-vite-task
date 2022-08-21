package user

import (
	"app/adapter/auth"
	"app/entity/model"
	"app/entity/model/db_model"
	"app/infra/database"
	"app/usecases"
	"time"
)

type CreateUserDependencies struct {
	gormHandler *database.GormHandler
}

func (dep CreateUserDependencies) Do(input model.InputUser) (*int, error) {
	currentTime := time.Now()

	createRecord := &db_model.User{
		UserName:  input.UserName,
		Email:     input.Email,
		UserIcon:  input.UserIcon,
		Password:  auth.HashMd5(input.Password),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	isOk := dep.gormHandler.DB.Create(createRecord).Error
	if isOk != nil {
		return nil, isOk
	}

	createId := int(createRecord.UserId)

	return &createId, nil
}

func NewCreateUserAdapter(gormHandler *database.GormHandler) usecases.CreateUserAdapter {
	return &CreateUserDependencies{gormHandler}
}
