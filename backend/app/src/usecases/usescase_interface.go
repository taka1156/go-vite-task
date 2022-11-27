package usecases

import "app/entity/model"

// user
type CreateUserInteface interface {
	Handle(input model.InputUser) (*model.User, error)
}

type GetUsersInterface interface {
	Handle() ([]*model.User, error)
}

type UpdateUserInteface interface {
	Handle(input model.UpdateUser) (*model.User, error)
}

type DeleteUserInterface interface {
	Handle(userId int) (bool, error)
}

// team
type CreateTeamInterface interface {
	Handle(input model.InputTeam) (*model.Team, error)
}

type GetTeamsInterface interface {
	Handle() ([]*model.Team, error)
}

type UpdateTeamInterface interface {
	Handle(input model.UpdateTeam) (*model.Team, error)
}

type DeleteTeamInterface interface {
	Handle(teamId int) (*bool, error)
}

// category
type CreateCategoryInterface interface {
	Handle(createInput model.InputCategory) (*model.Category, error)
}

type GetCategoriesInterface interface {
	Handle() ([]*model.Category, error)
}

type UpdateCategoryInterface interface {
	Handle(updateInput model.UpdateCategory) (*model.Category, error)
}

type DeleteCategoryInterface interface {
	Handle(categoryId int) (*bool, error)
}

// role
type CreateRoleInterface interface {
	Handle(input model.InputRole) (*model.Role, error)
}

type GetRoleInterface interface {
	Handle(roleId int) (*model.Role, error)
}

type GetRolesInterface interface {
	Handle() ([]*model.Role, error)
}

type UpdateRoleInterface interface {
	Handle(updateInput model.UpdateRole) (*model.Role, error)
}

type DeleteRoleInterface interface {
	Handle(roleId int) (*bool, error)
}
