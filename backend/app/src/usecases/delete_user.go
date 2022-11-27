package usecases

type DeleteUserAdapter interface {
	Do(userId int) (bool, error)
}

type DeleteUserDependencies struct {
	deleteUser DeleteUserAdapter
}

func (dep DeleteUserDependencies) Handle(userId int) (bool, error) {
	return dep.deleteUser.Do(userId)
}

func NewDeleteUserUsecase(deleteUser DeleteUserAdapter) DeleteUserInterface {
	return &DeleteUserDependencies{deleteUser}
}
