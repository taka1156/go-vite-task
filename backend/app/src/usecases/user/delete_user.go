package user

import "app/usecases"

type DeleteUserInteractor struct {
	userRepository usecases.UserRepository
}

func NewDeleteUserUsecase(userRepository usecases.UserRepository) usecases.DeleteUserUsecase {
	return &DeleteUserInteractor{userRepository}
}

func (interactor DeleteUserInteractor) Handle(userId int) (bool, error) {
	err := interactor.userRepository.DeleteUser(userId)
	if err != nil {
		return false, err
	}

	return true, nil
}
