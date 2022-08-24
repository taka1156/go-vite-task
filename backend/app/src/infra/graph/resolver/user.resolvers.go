package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/di"
	"app/entity/model"
	"context"
	"fmt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, inputUser model.InputUser) (*model.User, error) {
	user, err := di.Provider().CreateUser.Handle(inputUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, updateUser model.UpdateUser) (*model.User, error) {
	user, err := di.Provider().UpdateUser.Handle(updateUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, userID string, searchWord string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}
