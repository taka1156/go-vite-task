package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity/model"
	"app/infra/graph/generated"
	"context"
	"fmt"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, inputTask model.InputTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: CreateTask - createTask"))
}

// UpdateTask is the resolver for the updateTask field.
func (r *mutationResolver) UpdateTask(ctx context.Context, updateTask model.UpdateTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented: UpdateTask - updateTask"))
}

// DeleteTask is the resolver for the deleteTask field.
func (r *mutationResolver) DeleteTask(ctx context.Context, taskID *string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteTask - deleteTask"))
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, taskID string, searchWord string) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented: Task - task"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
