package user

import (
	"app/entity/convert"
	"app/entity/gmodel"
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"time"
)

type CreateUserDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep CreateUserDependencies) Do(input model.InputUser) (*int, error) {
	currentTime := time.Now()

	user := &gmodel.User{
		UserName:  input.UserName,
		Email:     input.Email,
		UserIcon:  input.UserIcon,
		Password:  input.Password,
		RoleID:    1,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	dep.gormAdapter.DB.Create(&user)

	insertId := convert.Id_Uint(user.UserID)

	return &insertId, nil
}

func NewCreateUserAdapter(gormAdapter *database.GormAdapter) usecases.CreateUserAdapter {
	return &CreateUserDependencies{gormAdapter}
}
