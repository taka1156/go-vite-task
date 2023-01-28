package repository

import (
	"app/entity/model"
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
)

type TaskStateRepositoryDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func NewTaskStateRepository(
	sqlAdapter *database.SqlAdapter,
) usecases.TaskStatusRepository {
	return &TaskStateRepositoryDependencies{sqlAdapter}
}

func (dep TaskStateRepositoryDependencies) GetTaskState(taskStateId int) (*model.TaskState, error) {

	taskState := model.TaskState{}

	return &taskState, nil
}

func (dep TaskStateRepositoryDependencies) GetTaskStatus() ([]*model.TaskState, error) {
	var taskStatus []*model.TaskState
	// stmt := `
	// select
	// 	us.user_id,
	// 	us.user_name,
	// 	us.email,
	// 	us.user_icon,
	// 	rs.role_id,
	// 	rs.role_name,
	// 	rs.role_icon
	// from
	// 	users as us
	// inner join
	// 	roles as rs
	// on
	// 	rs.role_id = us.role_id
	// where
	// 	users.user_id = ? and

	// 	user.deleted_at is null
	// `

	return taskStatus, nil
}

func (dep TaskStateRepositoryDependencies) CreateTaskState() error {
	// stmt := `
	// select
	// 	us.user_id,
	// 	us.user_name,
	// 	us.email,
	// 	us.user_icon,
	// 	rs.role_id,
	// 	rs.role_name,
	// 	rs.role_icon
	// from
	// 	users as us
	// inner join
	// 	roles as rs
	// on
	// 	rs.role_id = us.role_id
	// where
	// 	users.user_id = ? and

	// 	user.deleted_at is null
	// `

	return nil
}

func (dep TaskStateRepositoryDependencies) UpdateTaskState(input model.UpdateTaskState) (*int, error) {

	taskStatus := xo.TaskStatus{}

	return &taskStatus.TaskID, nil
}

func (dep TaskStateRepositoryDependencies) DeleteTaskState(taskStateId int) error {

	return nil
}
