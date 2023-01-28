package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type GetTeamsInteractorDependencies struct {
	teamRepository usecases.TeamRepository
}

func NewGetTeamsUsecase(teamRepository usecases.TeamRepository) usecases.GetTeamsUsecase {
	return &GetTeamsInteractorDependencies{teamRepository}
}

func (interactor GetTeamsInteractorDependencies) Handle() ([]*model.Team, error) {

	teams, err := interactor.teamRepository.GetTeams()
	if err != nil {
		return nil, err
	}
	return teams, nil
}
