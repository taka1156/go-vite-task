package usecases

import "app/entity/model"

type CreateUserAdapter interface {
	Do(user model.InputUser) (model.InputUser, error)
}

type CreateUserDpendences struct {
	createUser CreateUserAdapter
}

func (dep CreateUserDpendences) Hndle(user model.InputUser) (model.InputUser, error) {
	return dep.createUser.Do(user)
}

func NewCreateUserUsecase(createUser CreateUserAdapter) CreateUserInteface {
	return &CreateUserDpendences{createUser}
}
