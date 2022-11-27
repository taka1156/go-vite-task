package usecases

import (
	"app/entity/model"
)

type UpdateTeamAdapter interface {
	Do(input model.UpdateTeam) (*uint, error)
}

type UpdateTeamDependencies struct {
	updateTeam UpdateTeamAdapter
	getTeam    GetTeamAdapter
}

func (dep UpdateTeamDependencies) Handle(input model.UpdateTeam) (*model.Team, error) {
	teamId, err := dep.updateTeam.Do(input)
	if err != nil {
		return nil, err
	}

	team, err := dep.getTeam.Do(teamId)
	if err != nil {
		return nil, err
	}
	return team, nil
}

func NewUpdateTeamUsecase(updateTeam UpdateTeamAdapter, getTeam GetTeamAdapter) UpdateTeamInterface {
	return &UpdateTeamDependencies{updateTeam, getTeam}
}
