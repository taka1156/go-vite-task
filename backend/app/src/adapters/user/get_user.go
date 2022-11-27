package user

import (
	"app/entity/model"
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
	"context"
	"strconv"
)

type GetUserDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep GetUserDependencies) Do(userId *int) (*model.User, error) {

	user, err := xo.UserByUserID(context.TODO(), dep.sqlAdapter.DB, *userId)
	if err != nil {
		return nil, err
	}

	role, err := xo.RoleByRoleID(context.TODO(), dep.sqlAdapter.DB, int(user.RoleID.Int64))
	if err != nil {
		return nil, err
	}

	getUser := &model.User{
		UserID:   strconv.Itoa(user.UserID),
		UserName: user.UserName,
		UserIcon: &user.UserIcon.String,
		Email:    user.Email,
		Role: &model.Role{
			RoleID:   strconv.Itoa(role.RoleID),
			RoleName: role.RoleName,
			RoleIcon: &role.RoleIcon.String,
		},
	}

	return getUser, nil
}

func NewGetUserAdapter(sqlAdapter *database.SqlAdapter) usecases.GetUserAdapter {
	return &GetUserDependencies{sqlAdapter}
}
