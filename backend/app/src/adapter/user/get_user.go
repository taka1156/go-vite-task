package user

import (
	"app/entity/model"
	"app/infra/database"
	"app/usecases"
)

type GetUserDependencies struct {
	gormAdapter *database.GormAdapter
}

func (dep GetUserDependencies) Do(userId *int) (*model.User, error) {

	result := struct {
		UserID   string
		Email    string
		UserName string
		UserIcon *string
		RoleID   string
		RoleName string
		RoleIcon *string
	}{}

	dep.gormAdapter.DB.
		Table("users").
		Select(`users.user_id,
		        users.user_name,
				users.user_icon,
				users.email,
				roles.role_id,
				roles.role_name,
				roles.role_icon`,
		).
		Joins("LEFT OUTER JOIN roles ON roles.role_id = users.role_id").
		Where("users.user_id = ?", userId).Scan(&result)

	getUser := &model.User{
		UserID:   result.UserID,
		UserName: result.UserName,
		UserIcon: result.UserIcon,
		Email:    result.Email,
		Role: &model.Role{
			RoleID:   result.RoleID,
			RoleName: result.RoleName,
			RoleIcon: result.RoleIcon,
		},
	}

	return getUser, nil
}

func NewGetUserAdapter(gormAdapter *database.GormAdapter) usecases.GetUserAdapter {
	return &GetUserDependencies{gormAdapter}
}
