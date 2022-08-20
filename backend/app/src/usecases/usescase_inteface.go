package usecases

import "app/entity/model"

type CreateUserInteface interface {
	Hndle(user model.InputUser) (model.InputUser, error)
}
