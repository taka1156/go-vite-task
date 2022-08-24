package usecases

import (
	"app/entity/model"
)

type GetUsersAdapter interface {
	Do() ([]*model.User, error)
}

type GetUsersDependencies struct {
	getUsers GetUsersAdapter
}

func (dep GetUsersDependencies) Handle() ([]*model.User, error) {

	users, err := dep.getUsers.Do()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewGetUsersUsecase(getUsers GetUsersAdapter) GetUsersInterface {
	return &GetUsersDependencies{getUsers}
}
