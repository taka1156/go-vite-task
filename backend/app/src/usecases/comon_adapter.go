package usecases

import "app/entity/model"

type GetUserAdapter interface {
	Do(userId *int) (*model.User, error)
}

type GetTeamAdapter interface {
	Do(teamId *int) (*model.Team, error)
}
