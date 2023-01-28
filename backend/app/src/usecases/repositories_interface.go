package usecases

import "app/entity/model"

type RoleRepository interface {
	CanSelectRole(roleId int) error
}

type UserRepository interface {
	GetUser(userId int) (*model.User, error)
	GetUsers() ([]*model.User, error)
	CreateUser(input model.InputUser) (*int, error)
	UpdateUser(input model.UpdateUser) (*int, error)
	DeleteUser(userId int) error
}

type TeamRepository interface {
	GetTeam(teamId int) (*model.Team, error)
	GetTeams() ([]*model.Team, error)
	CreateTeam(input model.InputTeam) (*int, error)
	UpdateTeam(input model.UpdateTeam) (*int, error)
	DeleteTeam(teamId int) error
}

type TaskRepository interface {
	GetTask(taskId int) (*model.Task, error)
	GetTasks() ([]*model.Task, error)
	CreateTask() (*int, error)
	UpdateTask(input model.UpdateTask) (*int, error)
	DeleteTask(taskId int) error
}

type CategoryRepository interface{}

type TaskStatusRepository interface {
	GetTaskState(taskStateId int) (*model.TaskState, error)
	GetTaskStatus() ([]*model.TaskState, error)
	CreateTaskState() error
	UpdateTaskState(input model.UpdateTaskState) (*int, error)
	DeleteTaskState(taskStateId int) error
}

type TaskFlagRepository interface {
	CanSelectTaskFlag(taskId int)
}
