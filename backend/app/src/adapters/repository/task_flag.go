package repository

import (
	"app/infra/database"
	"app/usecases"
)

type TaskFlagRepositoryDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func NewTaskFlagRepository(
	sqlAdapter *database.SqlAdapter,
) usecases.TaskFlagRepository {
	return &TaskFlagRepositoryDependencies{sqlAdapter}
}

func (dep TaskFlagRepositoryDependencies) CanSelectTaskFlag(taskId int) {
	// taskFlagIdが存在するか確認
}
