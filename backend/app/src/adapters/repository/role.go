package repository

import (
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
	"errors"
)

type RoleRepositoryDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func NewRoleRepository(
	sqlAdapter *database.SqlAdapter,
) usecases.RoleRepository {
	return &RoleRepositoryDependencies{sqlAdapter}
}

func (dep RoleRepositoryDependencies) CanSelectRole(roleId int) error {
	// roleIdが存在するか確認

	role := &xo.Role{
		RoleID: roleId,
	}

	existsRoleId := role.Exists()
	if !existsRoleId {
		return errors.New("Invalid value for roleId.")
	}

	return nil
}
