package usecases

type DeleteTeamAdapter interface {
	Do(userId int) (*bool, error)
}

type DeleteTeamDependencies struct {
	deleteTeam DeleteTeamAdapter
}

func (dep DeleteTeamDependencies) Handle(userId int) (*bool, error) {
	isDeleted, err := dep.deleteTeam.Do(userId)
	if err != nil {
		return nil, err
	}

	return isDeleted, nil
}

func NewDeleteTeamUsecase(deleteTeam DeleteTeamAdapter) DeleteTeamInterface {
	return &DeleteTeamDependencies{deleteTeam}
}
