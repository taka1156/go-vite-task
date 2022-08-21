package usecases

import "app/entity/model"

type CreateUserInteface interface {
	Handle(input model.InputUser) (*model.User, error)
}

type UpdateUserInteface interface {
	Handle(input model.UpdateUser) (*model.User, error)
}
