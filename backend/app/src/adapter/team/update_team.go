package team

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
	"time"
)

type UpdateTeamDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep UpdateTeamDependencies) Do(input model.UpdateTeam) (*int, error) {
	currentTime := time.Now()

	return &updateTeamId, nil
}

func NewUpdateTeamAdapter(sqlAdapter *database.SqlAdapter) usecases.UpdateTeamAdapter {
	return &UpdateTeamDependencies{sqlAdapter}
}
