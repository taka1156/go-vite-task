package usecases

import (
	"app/entity/model"
	"app/usecases"
)

type CreateRoleInteractorDependencies struct {
	roleRepository usecases.RoleRepository
}

func NewCreateRoleUsecase(roleRepository usecases.RoleRepository) usecases.CreateRoleUsecase {
	return &CreateRoleInteractorDependencies{roleRepository}
}

func (interactor CreateRoleInteractorDependencies) Handle(inputRole model.InputRole) (*model.Role, error) {
	// roleId, err := interactor.roleRepository.CreateRole(inputRole)
	// if err != nil {
	// 	return nil, err
	// }

	// user, err := interactor.roleRepository.GetRole(roleId)
	// if err != nil {
	// 	return nil, err
	// }

	role := model.Role{}

	return &role, nil
}
