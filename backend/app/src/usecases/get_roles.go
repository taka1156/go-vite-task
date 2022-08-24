package usecases

import "app/entity/model"

type GetRolesAdapter interface {
	Do() ([]*model.Role, error)
}

type GetRolesDependencies struct {
	getRoles GetRolesAdapter
}

func (dep GetRolesDependencies) Handle() ([]*model.Role, error) {
	roles, err := dep.getRoles.Do()
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func NewGetRolesUsecase(getRoles GetRolesAdapter) GetRolesInterface {
	return &GetRolesDependencies{getRoles}
}
