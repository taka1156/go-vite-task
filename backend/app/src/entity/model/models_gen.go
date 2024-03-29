// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	CategoryID   string  `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	CategoryIcon *string `json:"categoryIcon"`
}

type InputCategory struct {
	CategoryName string  `json:"categoryName"`
	CategoryIcon *string `json:"categoryIcon"`
}

type InputRole struct {
	RoleName string  `json:"roleName"`
	RoleIcon *string `json:"roleIcon"`
}

type InputTask struct {
	TaskName    string  `json:"taskName"`
	Link        *string `json:"link"`
	Description *string `json:"description"`
}

type InputTaskState struct {
	Progress     int     `json:"progress"`
	TaskID       *string `json:"taskId"`
	TaskName     string  `json:"taskName"`
	Link         *string `json:"link"`
	Description  *string `json:"description"`
	TaskFlagID   string  `json:"taskFlagId"`
	TaskFlagName string  `json:"taskFlagName"`
}

type InputTeam struct {
	TeamName  string   `json:"teamName"`
	TeamIcon  *string  `json:"teamIcon"`
	TeamMemo  *string  `json:"teamMemo"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Users     []string `json:"users"`
}

type InputUser struct {
	UserName string  `json:"userName"`
	Email    string  `json:"email"`
	UserIcon *string `json:"userIcon"`
	Password string  `json:"password"`
}

type Role struct {
	RoleID   string  `json:"roleId"`
	RoleName string  `json:"roleName"`
	RoleIcon *string `json:"roleIcon"`
}

type Task struct {
	TaskID      *string `json:"taskId"`
	TaskName    string  `json:"taskName"`
	Link        *string `json:"link"`
	Description *string `json:"description"`
}

type TaskFlag struct {
	TaskFlagID   string `json:"taskFlagId"`
	TaskFlagName string `json:"taskFlagName"`
}

type TaskState struct {
	Task     *Task     `json:"task"`
	TaskFlag *TaskFlag `json:"taskFlag"`
	Progress int       `json:"progress"`
}

type Team struct {
	TeamID    string  `json:"teamId"`
	TeamName  string  `json:"teamName"`
	TeamIcon  *string `json:"teamIcon"`
	TeamMemo  *string `json:"teamMemo"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Users     []*User `json:"users"`
}

type UpdateCategory struct {
	CategoryID   string  `json:"categoryId"`
	CategoryName string  `json:"categoryName"`
	CategoryIcon *string `json:"categoryIcon"`
}

type UpdateRole struct {
	RoleID   string  `json:"roleId"`
	RoleName string  `json:"roleName"`
	RoleIcon *string `json:"roleIcon"`
}

type UpdateTask struct {
	TaskID      *string `json:"taskId"`
	TaskName    string  `json:"taskName"`
	Link        *string `json:"link"`
	Description *string `json:"description"`
}

type UpdateTaskState struct {
	TaskStateID  *string `json:"taskStateId"`
	Progress     int     `json:"progress"`
	TaskID       *string `json:"taskId"`
	TaskName     string  `json:"taskName"`
	Link         *string `json:"link"`
	Description  *string `json:"description"`
	TaskFlagID   string  `json:"taskFlagId"`
	TaskFlagName string  `json:"taskFlagName"`
}

type UpdateTeam struct {
	TeamID    string   `json:"teamId"`
	TeamName  string   `json:"teamName"`
	TeamIcon  *string  `json:"teamIcon"`
	TeamMemo  *string  `json:"teamMemo"`
	StartDate string   `json:"startDate"`
	EndDate   string   `json:"endDate"`
	Users     []string `json:"users"`
}

type UpdateUser struct {
	UserID   string  `json:"userId"`
	UserName string  `json:"userName"`
	Email    string  `json:"email"`
	UserIcon *string `json:"userIcon"`
	Password string  `json:"password"`
	RoleID   *string `json:"roleId"`
}

type User struct {
	UserID   string  `json:"userId"`
	UserName string  `json:"userName"`
	UserIcon *string `json:"userIcon"`
	Email    string  `json:"email"`
	Role     *Role   `json:"role"`
}
