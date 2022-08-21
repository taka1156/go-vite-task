package user

import (
	"app/entity/model"
	"app/entity/model/db_model"
	"app/infra/database"
	"app/usecases"
	"strconv"
	"time"
)

type UpdateUserDependencies struct {
	gormHandler *database.GormHandler
}

func (dep UpdateUserDependencies) Do(input model.UpdateUser) (*int, error) {
	currentTime := time.Now()

	updateRecord := &db_model.User{
		UserName:  input.UserName,
		Email:     input.Email,
		UserIcon:  input.UserIcon,
		UpdatedAt: currentTime,
	}

	isOk := dep.gormHandler.DB.Model(&db_model.User{}).Where("user_id = ?", input.UserID).Updates(updateRecord).Error
	if isOk != nil {
		return nil, isOk
	}

	updateId, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, err
	}

	return &updateId, nil
}

func NewUpdateUserAdapter(gormHandler *database.GormHandler) usecases.UpdateUserAdapter {
	return &UpdateUserDependencies{gormHandler}
}
