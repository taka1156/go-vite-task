package di

import (
	"app/adapters/user"
	"app/envvars"
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

	// Env
	envvars.EnvLoad()

	// Infrastructure
	gormAdapter := database.GormInit(envvars.DbEnv)

	// Adapter
	createUserAdapter := user.NewCreateUserAdapter(gormAdapter)
	getUserAdapter := user.NewGetUserAdapter(gormAdapter)
	updateUserAdapter := user.NewUpdateUserAdapter(gormAdapter)

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
