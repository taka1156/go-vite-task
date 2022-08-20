package user

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type CreateUserDpendences struct {
	gormAdapter database.gormAdapter
}

func (dep CreateUserDpendences) Do(input model.InputUser) (model.InputUser, error) {

	return input, nil
}

func NewCreateUserAdapter(gormAdapter database.gormAdapter) usecases.CreateUserAdapter {
	return &CreateUserDpendences{gormAdapter}
}
