package usecases

import "app/entity/model"

// user
type GetUserAdapter interface {
	Do(userId *int) (*model.User, error)
}

// team
type GetTeamAdapter interface {
	Do(teamId *int) (*model.Team, error)
}

type GetTeamUsersAdapter interface {
	Do(teamId *int) ([]*model.User, error)
}

// category
type GetCategoryAdapter interface {
	Do(teamId *int) (*model.Category, error)
}

// role
type GetRoleAdapter interface {
	Do(teamId *int) (*model.Role, error)
}
