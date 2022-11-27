package usecases

import "app/entity/model"

// user
type GetUserAdapter interface {
	Do(userId *int) (*model.User, error)
}

// team
type GetTeamAdapter interface {
	Do(teamId *uint) (*model.Team, error)
}

type GetTeamUsersAdapter interface {
	Do(teamId *uint) ([]*model.User, error)
}

// category
type GetCategoryAdapter interface {
	Do(teamId *uint) (*model.Category, error)
}

// role
type GetRoleAdapter interface {
	Do(teamId *uint) (*model.Role, error)
}
