package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type CreateTeamInteractorDependencies struct {
	teamRepository usecases.TeamRepository
}

func NewCreateTeamUsecase(teamRepository usecases.TeamRepository) usecases.CreateTeamUsecase {
	return &CreateTeamInteractorDependencies{teamRepository}
}

func (interactor CreateTeamInteractorDependencies) Handle(input model.InputTeam) (*model.Team, error) {
	teamId, err := interactor.teamRepository.CreateTeam(input)
	if err != nil {
		return nil, err
	}

	user, err := interactor.teamRepository.GetTeam(*teamId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
