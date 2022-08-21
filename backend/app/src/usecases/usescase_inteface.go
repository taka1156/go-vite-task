package usecases

import "app/entity/model"

type CreateUserInteface interface {
	Handle(input model.InputUser) (*model.User, error)
}

type UpdateUserInteface interface {
	Handle(input model.UpdateUser) (*model.User, error)
}

type CreateTeamInterface interface {
	Handle(input model.InputTeam) (*model.Team, error)
}

type UpdateTeamInteface interface {
	Handle(input model.UpdateTeam) (*model.Team, error)
}
