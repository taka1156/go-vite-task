package usecases

import (
	"app/entity/model"
)

type UpdateUserAdapter interface {
	Do(user model.UpdateUser) (*int, error)
}

type UpdateUserDependencies struct {
	updateUser UpdateUserAdapter
	getUser    GetUserAdapter
}

func (dep UpdateUserDependencies) Handle(input model.UpdateUser) (*model.User, error) {
	userId, err := dep.updateUser.Do(input)
	if err != nil {
		return nil, err
	}

	user, err := dep.getUser.Do(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewUpdateUserUsecase(updateUser UpdateUserAdapter, getUser GetUserAdapter) UpdateUserInteface {
	return &UpdateUserDependencies{updateUser, getUser}
}
