package user

import (
	"app/entity/model"
	"app/usecases"
)

type UpdateUserInteractor struct {
	userRepository usecases.UserRepository
}

func NewUpdateUserUsecase(userRepository usecases.UserRepository) usecases.UpdateUserUsecase {
	return &UpdateUserInteractor{userRepository}
}

func (interactor UpdateUserInteractor) Handle(input model.UpdateUser) (*model.User, error) {
	userId, err := interactor.userRepository.UpdateUser(input)
	if err != nil {
		return nil, err
	}

	user, err := interactor.userRepository.GetUser(*userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
