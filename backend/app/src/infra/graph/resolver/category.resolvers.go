package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity/model"
	"app/infra/graph/generated"
	"context"
	"fmt"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, inputCategory model.InputCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: CreateCategory - createCategory"))
}

// UpdateCategory is the resolver for the updateCategory field.
func (r *mutationResolver) UpdateCategory(ctx context.Context, updateCategory model.UpdateCategory) (*model.Category, error) {
	panic(fmt.Errorf("not implemented: UpdateCategory - updateCategory"))
}

// DeleteCategory is the resolver for the deleteCategory field.
func (r *mutationResolver) DeleteCategory(ctx context.Context, categoryID string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteCategory - deleteCategory"))
}

// Category is the resolver for the category field.
func (r *queryResolver) Category(ctx context.Context, categoryID string, searchWord string) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented: Category - category"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
