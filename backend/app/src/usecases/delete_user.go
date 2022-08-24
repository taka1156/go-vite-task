package usecases

type DeleteUserAdapter interface {
	Do(userId int) (*bool, error)
}

type DeleteUserDependencies struct {
	deleteUser DeleteUserAdapter
}

func (dep DeleteUserDependencies) Handle(userId int) (*bool, error) {
	isDeleted, err := dep.deleteUser.Do(userId)
	if err != nil {
		return nil, err
	}

	return isDeleted, nil
}

func NewDeleteUserUsecase(deleteUser DeleteUserAdapter) DeleteUserInterface {
	return &DeleteTeamDependencies{deleteUser}
}
