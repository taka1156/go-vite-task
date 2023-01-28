package user

import (
	"app/entity/model"
	"app/usecases"
)

type GetUsersInteractorDependencies struct {
	userRepository usecases.UserRepository
}

func NewGetUsersUsecase(userRepository usecases.UserRepository) usecases.GetUsersUsecase {
	return &GetUsersInteractorDependencies{userRepository}
}

func (interactor GetUsersInteractorDependencies) Handle() ([]*model.User, error) {
	users, err := interactor.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
