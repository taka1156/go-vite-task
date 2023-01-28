package di

import (
	"app/adapters/repository"
	"app/envvars"
	"app/infra/database"
	"app/usecases"
	"app/usecases/user"
)

type Function struct {
	CreateUser usecases.CreateUserUsecase
	UpdateUser usecases.UpdateUserUsecase
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
	dbAdapter := database.DBInit(envvars.DbEnv)

	// repository
	userRepository := repository.NewUserRepository(dbAdapter)

	// Usecase
	createUserUsecase := user.NewCreateUserUsecase(userRepository)
	updateUserUsecase := user.NewUpdateUserUsecase(userRepository)

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
