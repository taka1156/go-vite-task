package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type UpdateTeamInteractorDependencies struct {
	teamRepository usecases.TeamRepository
}

func NewUpdateTeamUsecase(teamRepository usecases.TeamRepository) usecases.UpdateTeamUsecase {
	return &UpdateTeamInteractorDependencies{teamRepository}
}

func (interactor UpdateTeamInteractorDependencies) Handle(input model.UpdateTeam) (*model.Team, error) {
	teamId, err := interactor.teamRepository.UpdateTeam(input)
	if err != nil {
		return nil, err
	}

	team, err := interactor.teamRepository.GetTeam(*teamId)
	if err != nil {
		return nil, err
	}
	return team, nil
}
