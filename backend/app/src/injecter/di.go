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
	sqlAdapter := database.DbInit()

	// Adapter
	createUserAdapter := user.NewCreateUserAdapter(sqlAdapter)
	getUserAdapter := user.NewGetUserAdapter(sqlAdapter)
	updateUserAdapter := user.NewUpdateUserAdapter(sqlAdapter)

	// Usecase
	creatUserUsecase := usecases.NewCreateUserUsecase(createUserAdapter, getUserAdapter)
	updateUserUsecase := usecases.NewUpdateUserUsecase(updateUserAdapter, getUserAdapter)

	f := &Function{
		createUser: creatUserUsecase,
		updateUser: updateUserUsecase,
	}

	Di.dep = *f

}
