package usecases

type DeleteRoleAdapter interface {
	Do(roleId int) (*bool, error)
}

type DeleteRoleDependencies struct {
	deleteRole DeleteRoleAdapter
}

func (dep DeleteRoleDependencies) Handle(roleId int) (*bool, error) {
	isDeleted, err := dep.deleteRole.Do(roleId)
	if err != nil {
		return nil, err
	}

	return isDeleted, nil
}

func NewDeleteRoleUsecase(deleteRole DeleteRoleAdapter) DeleteRoleInterface {
	return &DeleteRoleDependencies{deleteRole}
}
