package user

import (
	"app/entity/model"
	"app/usecases"
)

type CreateUserInteractor struct {
	userRepository usecases.UserRepository
}

func NewCreateUserUsecase(userRepository usecases.UserRepository) usecases.CreateUserUsecase {
	return &CreateUserInteractor{userRepository}
}

func (interactor CreateUserInteractor) Handle(input model.InputUser) (*model.User, error) {
	userId, err := interactor.userRepository.CreateUser(input)
	if err != nil {
		return nil, err
	}

	user, err := interactor.userRepository.GetUser(*userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
