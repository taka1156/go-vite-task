package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type GetTeamsDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep GetTeamsDependencies) Do() ([]*model.Team, error) {
	var teams []*model.Team
	//currentTime := time.Now()

	return teams, nil
}

func NewGetTeamsAdapter(gormAdapter *database.GormAdapter) usecases.GetTeamsAdapter {
	return &GetTeamsDependencies{gormAdapter}
}
