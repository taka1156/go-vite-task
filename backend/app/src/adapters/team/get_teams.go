package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type GetTeamsDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep GetTeamsDependencies) Do() ([]*model.Team, error) {
	var teams []*model.Team
	//currentTime := time.Now()

	return teams, nil
}

func NewGetTeamsAdapter(sqlAdapter *database.SqlAdapter) usecases.GetTeamsAdapter {
	return &GetTeamsDependencies{sqlAdapter}
}
