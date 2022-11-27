package user

import (
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
	"context"
)

type DeleteUserDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep DeleteUserDependencies) Do(userId int) (bool, error) {

	user := xo.User{
		UserID: userId,
	}

	err := user.Delete(context.TODO(), dep.sqlAdapter.DB)
	if err != nil {
		return false, err
	}

	isDeleted := true

	return isDeleted, nil
}

func NewDeleteUserAdapter(sqlAdapter *database.SqlAdapter) usecases.DeleteUserAdapter {
	return &DeleteUserDependencies{sqlAdapter}
}
