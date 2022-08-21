package injecter

import (
	"app/adapter/user"
	"app/infra/database"
	"app/usecases"
)

var Di DiContainer

type DiContainer struct {
	dep Function
}

type Function struct {
	createUser usecases.CreateUserInteface
	updateUser usecases.UpdateUserInteface
}

func Injection() {

	// Infrastructure
	gormHandler := database.DbInit()

	// Adapter
	createUserAdapter := user.NewCreateUserAdapter(gormHandler)
	getUserAdapter := user.NewGetUserAdapter(gormHandler)
	updateUserAdapter := user.NewUpdateUserAdapter(gormHandler)

	// Usecase
	creatUserUsecase := usecases.NewCreateUserUsecase(createUserAdapter, getUserAdapter)
	updateUserUsecase := usecases.NewUpdateUserUsecase(updateUserAdapter, getUserAdapter)

	f := &Function{
		createUser: creatUserUsecase,
		updateUser: updateUserUsecase,
	}

	Di.dep = *f

}
