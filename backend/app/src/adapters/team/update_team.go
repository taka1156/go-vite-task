package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type UpdateTeamDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep UpdateTeamDependencies) Do(input model.UpdateTeam) (*uint, error) {
	//currentTime := time.Now()

	updateTeamId := uint(1)

	return &updateTeamId, nil
}

func NewUpdateTeamAdapter(sqlAdapter *database.SqlAdapter) usecases.UpdateTeamAdapter {
	return &UpdateTeamDependencies{sqlAdapter}
}
