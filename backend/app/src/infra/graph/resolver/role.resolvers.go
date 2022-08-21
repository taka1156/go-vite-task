package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity/model"
	"app/infra/graph/generated"
	"context"
	"fmt"
)

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, inputRole model.InputRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: CreateRole - createRole"))
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, updateRole model.UpdateRole) (*model.Role, error) {
	panic(fmt.Errorf("not implemented: UpdateRole - updateRole"))
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, roleID string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteRole - deleteRole"))
}

// Role is the resolver for the role field.
func (r *queryResolver) Role(ctx context.Context, roleID string, searchWord string) ([]*model.Role, error) {
	panic(fmt.Errorf("not implemented: Role - role"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
