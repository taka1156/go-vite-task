package usecases

import "app/entity/model"

// user
type CreateUserUsecase interface {
	Handle(input model.InputUser) (*model.User, error)
}

type GetUsersUsecase interface {
	Handle() ([]*model.User, error)
}

type UpdateUserUsecase interface {
	Handle(input model.UpdateUser) (*model.User, error)
}

type DeleteUserUsecase interface {
	Handle(userId int) (bool, error)
}

// team
type CreateTeamUsecase interface {
	Handle(input model.InputTeam) (*model.Team, error)
}

type GetTeamsUsecase interface {
	Handle() ([]*model.Team, error)
}

type UpdateTeamUsecase interface {
	Handle(input model.UpdateTeam) (*model.Team, error)
}

type DeleteTeamUsecase interface {
	Handle(teamId int) (bool, error)
}

// category
type CreateCategoryUsecase interface {
	Handle(createInput model.InputCategory) (*model.Category, error)
}

type GetCategoriesUsecase interface {
	Handle() ([]*model.Category, error)
}

type UpdateCategoryUsecase interface {
	Handle(updateInput model.UpdateCategory) (*model.Category, error)
}

type DeleteCategoryUsecase interface {
	Handle(categoryId int) (bool, error)
}

// role
type CreateRoleUsecase interface {
	Handle(input model.InputRole) (*model.Role, error)
}

type GetRoleUsecase interface {
	Handle(roleId int) (*model.Role, error)
}

type GetRolesUsecase interface {
	Handle() ([]*model.Role, error)
}

type UpdateRoleUsecase interface {
	Handle(updateInput model.UpdateRole) (*model.Role, error)
}

type DeleteRoleUsecase interface {
	Handle(roleId int) (bool, error)
}
