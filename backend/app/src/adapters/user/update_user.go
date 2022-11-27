package user

import (
	"app/entity/model"
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
	"context"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type UpdateUserDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func (dep UpdateUserDependencies) canSelectRole(roleId string) error {
	// roleIdが存在するか確認

	castRoleId, err := strconv.Atoi(roleId)
	if err != nil {
		return err
	}

	role := &xo.Role{
		RoleID: castRoleId,
	}

	existsRoleId := role.Exists()
	if !existsRoleId {
		return errors.New("Invalid value for roleId.")
	}

	return nil
}

func (dep UpdateUserDependencies) Do(input model.UpdateUser) (*int, error) {
	if input.RoleID != nil {
		if err := dep.canSelectRole(*input.RoleID); err != nil {
			return nil, err
		}
	}

	currentTime := time.Now()

	castUserID, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, err
	}
	castRoleID, err := strconv.Atoi(*input.RoleID)
	if err != nil {
		return nil, err
	}

	user := xo.User{
		UserID:    castUserID,
		UserName:  input.UserName,
		UserIcon:  sql.NullString{String: *input.UserIcon},
		Email:     input.Email,
		Password:  input.Password,
		RoleID:    sql.NullInt64{Int64: int64(castRoleID)},
		UpdatedAt: sql.NullTime{Time: currentTime},
	}

	user.Update(context.TODO(), dep.sqlAdapter.DB)

	return &user.UserID, nil
}

func NewUpdateUserAdapter(sqlAdapter *database.SqlAdapter) usecases.UpdateUserAdapter {
	return &UpdateUserDependencies{sqlAdapter}
}
