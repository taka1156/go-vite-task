package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type GetRolesInteractorDependencies struct {
	roleRepository usecases.RoleRepository
}

func NewGetRolesUsecase(rolesRepository usecases.RoleRepository) usecases.GetRolesUsecase {
	return &GetRolesInteractorDependencies{rolesRepository}
}

func (interactor GetRolesInteractorDependencies) Handle() ([]*model.Role, error) {
	roles := []*model.Role{}

	return roles, nil
}
