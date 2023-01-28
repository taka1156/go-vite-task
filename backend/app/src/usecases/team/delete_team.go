package usecases

import "app/usecases"

type DeleteTeamInteractorDependencies struct {
	teamRepository usecases.TeamRepository
}

func NewDeleteTeamUsecase(teamRepository usecases.TeamRepository) usecases.DeleteTeamUsecase {
	return &DeleteTeamInteractorDependencies{teamRepository}
}

func (interactor DeleteTeamInteractorDependencies) Handle(userId int) (bool, error) {
	err := interactor.teamRepository.DeleteTeam(userId)
	if err != nil {
		return false, err
	}

	return true, nil
}
