package usecases

import "app/entity/model"

type CreateUserAdapter interface {
	Do(user model.InputUser) (*int, error)
}

type CreateUserDependencies struct {
	createUser CreateUserAdapter
	getUser    GetUserAdapter
}

func (dep CreateUserDependencies) Handle(input model.InputUser) (*model.User, error) {
	userId, err := dep.createUser.Do(input)
	if err != nil {
		return nil, err
	}

	user, err := dep.getUser.Do(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewCreateUserUsecase(createUser CreateUserAdapter, getUser GetUserAdapter) CreateUserInteface {
	return &CreateUserDependencies{createUser, getUser}
}
