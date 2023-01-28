package repository

import (
	"app/entity/model"
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
)

type TaskRepositoryDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func NewTaskRepositoryDependencies(
	sqlAdapter *database.SqlAdapter,
) usecases.TaskRepository {
	return &TaskRepositoryDependencies{sqlAdapter}
}

func (dep TaskRepositoryDependencies) GetTask(taskId int) (*model.Task, error) {

	task := model.Task{}

	return &task, nil
}

func (dep TaskRepositoryDependencies) GetTasks() ([]*model.Task, error) {
	var tasks []*model.Task
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

	return tasks, nil
}

func (dep TaskRepositoryDependencies) CreateTask() (*int, error) {
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

	task := xo.Task{}

	return &task.TaskID, nil
}

func (dep TaskRepositoryDependencies) UpdateTask(input model.UpdateTask) (*int, error) {

	updateTask := xo.Task{}

	return &updateTask.TaskID, nil
}

func (dep TaskRepositoryDependencies) DeleteTask(taskId int) error {

	return nil
}

func (dep TaskRepositoryDependencies) CanSelectTask(taskId string) (*bool, error) {
	// taskIdが存在するか確認

	isSelect := false

	return &isSelect, nil
}
