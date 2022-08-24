package usecases

import (
	"app/entity/model"
)

type GetTeamsAdapter interface {
	Do() ([]*model.Team, error)
}

type GetTeamsDependencies struct {
	getTeams GetTeamsAdapter
}

func (dep GetTeamsDependencies) Handle() ([]*model.Team, error) {

	teams, err := dep.getTeams.Do()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func NewGetTeamsUsecase(getTeams GetTeamsAdapter) GetTeamsInterface {
	return &GetTeamsDependencies{getTeams}
}
