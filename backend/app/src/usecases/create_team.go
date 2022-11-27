package usecases

import "app/entity/model"

type CreateTeamAdapter interface {
	Do(user model.InputTeam) (*uint, error)
}

type CreateTeamDependencies struct {
	createTeam CreateTeamAdapter
	getTeam    GetTeamAdapter
}

func (dep CreateTeamDependencies) Handle(input model.InputTeam) (*model.Team, error) {
	teamId, err := dep.createTeam.Do(input)
	if err != nil {
		return nil, err
	}

	user, err := dep.getTeam.Do(teamId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewCreateTeamUsecase(createTeam CreateTeamAdapter, getTeam GetTeamAdapter) CreateTeamInterface {
	return &CreateTeamDependencies{createTeam, getTeam}
}
