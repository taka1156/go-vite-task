package usecases

import "app/usecases"

type DeleteRoleInteractorDependencies struct {
	roleRepository usecases.RoleRepository
}

func NewDeleteRoleUsecase(roleRepository usecases.RoleRepository) usecases.DeleteRoleUsecase {
	return &DeleteRoleInteractorDependencies{roleRepository}
}

func (interactor DeleteRoleInteractorDependencies) Handle(roleId int) (bool, error) {
	// err := interactor.roleRepository.DeleteRole(roleId)
	// if err != nil {
	// 	return nil, err
	// }

	return true, nil
}
