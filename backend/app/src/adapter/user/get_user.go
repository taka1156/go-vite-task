package user

import (
	"app/entity/model"
	"app/entity/model/db_model"
	"app/infra/database"
	"app/usecases"
	"strconv"
)

type GetUserDependencies struct {
	gormHandler *database.GormHandler
}

func (dep GetUserDependencies) Do(userId *int) (*model.User, error) {
	getReacord := &db_model.User{}

	err := dep.gormHandler.DB.Where("user_id = ?", userId).First(getReacord)
	if err != nil {
		return nil, err.Error
	}

	resultRecord := &model.User{
		UserID:   strconv.Itoa(int(getReacord.UserId)),
		UserName: getReacord.UserName,
		UserIcon: getReacord.UserIcon,
		Email:    getReacord.Email,
		RoleID:   nil,
	}

	return resultRecord, nil
}

func NewGetUserAdapter(gormHandler *database.GormHandler) usecases.GetUserAdapter {
	return &GetUserDependencies{gormHandler}
}
