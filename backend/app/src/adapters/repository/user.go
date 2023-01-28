package repository

import (
	"app/entity/model"
	"app/entity/model/xo"
	"app/infra/database"
	"app/usecases"
	"context"
	"database/sql"
	"strconv"
	"time"
)

type UserRepositoryDependencies struct {
	sqlAdapter *database.SqlAdapter
}

func NewUserRepository(
	sqlAdapter *database.SqlAdapter,
) usecases.UserRepository {
	return &UserRepositoryDependencies{sqlAdapter}
}

func (dep UserRepositoryDependencies) GetUser(userId int) (*model.User, error) {

	user, err := xo.UserByUserID(context.TODO(), dep.sqlAdapter.DB, userId)
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

func (dep UserRepositoryDependencies) GetUsers() ([]*model.User, error) {
	var users []*model.User
	// stmt := `
	// select
	// 	us.user_id,
	// 	us.user_name,
	// 	us.email,
	// 	us.user_icon,
	// 	rs.role_id,
	// 	rs.role_name,
	// 	rs.role_icon
	// from
	// 	users as us
	// inner join
	// 	roles as rs
	// on
	// 	rs.role_id = us.role_id
	// where
	// 	users.user_id = ? and

	// 	user.deleted_at is null
	// `

	return users, nil
}

func (dep UserRepositoryDependencies) CreateUser(input model.InputUser) (*int, error) {
	// stmt := `
	// select
	// 	us.user_id,
	// 	us.user_name,
	// 	us.email,
	// 	us.user_icon,
	// 	rs.role_id,
	// 	rs.role_name,
	// 	rs.role_icon
	// from
	// 	users as us
	// inner join
	// 	roles as rs
	// on
	// 	rs.role_id = us.role_id
	// where
	// 	users.user_id = ? and

	// 	user.deleted_at is null
	// `

	user := xo.User{}

	return &user.UserID, nil
}

func (dep UserRepositoryDependencies) UpdateUser(input model.UpdateUser) (*int, error) {

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

func (dep UserRepositoryDependencies) DeleteUser(userId int) error {

	user := xo.User{
		UserID: userId,
	}

	err := user.Delete(context.TODO(), dep.sqlAdapter.DB)
	if err != nil {
		return err
	}

	return nil
}
