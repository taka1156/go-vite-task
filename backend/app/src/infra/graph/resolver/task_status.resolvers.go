package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity/model"
	"context"
	"fmt"
)

// CreateTaskState is the resolver for the createTaskState field.
func (r *mutationResolver) CreateTaskState(ctx context.Context, inputTaskState model.InputTaskState) (*model.TaskState, error) {
	panic(fmt.Errorf("not implemented: CreateTaskState - createTaskState"))
}

// UpdateTaskState is the resolver for the updateTaskState field.
func (r *mutationResolver) UpdateTaskState(ctx context.Context, updateTaskState model.UpdateTaskState) (*model.TaskState, error) {
	panic(fmt.Errorf("not implemented: UpdateTaskState - updateTaskState"))
}

// DeleteTaskState is the resolver for the deleteTaskState field.
func (r *mutationResolver) DeleteTaskState(ctx context.Context, taskStateID string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteTaskState - deleteTaskState"))
}

// TaskState is the resolver for the taskState field.
func (r *queryResolver) TaskState(ctx context.Context, taskID string, searchWord string) ([]*model.TaskState, error) {
	panic(fmt.Errorf("not implemented: TaskState - taskState"))
}
