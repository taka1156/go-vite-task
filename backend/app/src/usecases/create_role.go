package usecases

import "app/entity/model"

type CreateRoleAdapter interface {
	Do(inputRole model.InputRole) (*uint, error)
}

type CreateRoleDependencies struct {
	createRole CreateRoleAdapter
	getRole    GetRoleAdapter
}

func (dep CreateRoleDependencies) Handle(inputRole model.InputRole) (*model.Role, error) {
	roleId, err := dep.createRole.Do(inputRole)
	if err != nil {
		return nil, err
	}

	user, err := dep.getRole.Do(roleId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewCreateRoleUsecase(createRole CreateRoleAdapter, getRole GetRoleAdapter) CreateRoleInterface {
	return &CreateRoleDependencies{createRole, getRole}
}
