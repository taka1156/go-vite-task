package di

import (
	"app/adapter/user"
	"app/infra/database"
	"app/usecases"
)

type Function struct {
	CreateUser usecases.CreateUserInteface
	UpdateUser usecases.UpdateUserInteface
}

type Middleware struct {
}

type Adapter struct {
}

var function Function
var middleware Middleware
var adapter Adapter

func Do() {

	// Infrastructure
	sqlAdapter := database.DbInit()

	// Adapter
	createUserAdapter := user.NewCreateUserAdapter(sqlAdapter)
	getUserAdapter := user.NewGetUserAdapter(sqlAdapter)
	updateUserAdapter := user.NewUpdateUserAdapter(sqlAdapter)

	// Usecase
	createUserUsecase := usecases.NewCreateUserUsecase(createUserAdapter, getUserAdapter)
	updateUserUsecase := usecases.NewUpdateUserUsecase(updateUserAdapter, getUserAdapter)

	function = Function{
		CreateUser: createUserUsecase,
		UpdateUser: updateUserUsecase,
	}

}

func Provider() Function {
	return function
}

func ProviderMiddleware() Middleware {
	return middleware
}
