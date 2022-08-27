package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type UpdateTeamDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep UpdateTeamDependencies) Do(input model.UpdateTeam) (*int, error) {
	//currentTime := time.Now()

	updateTeamId := 1

	return &updateTeamId, nil
}

func NewUpdateTeamAdapter(gormAdapter *database.GormAdapter) usecases.UpdateTeamAdapter {
	return &UpdateTeamDependencies{gormAdapter}
}
