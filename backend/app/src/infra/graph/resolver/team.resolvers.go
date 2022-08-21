package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/entity/model"
	"context"
	"fmt"
)

// CreateTeam is the resolver for the createTeam field.
func (r *mutationResolver) CreateTeam(ctx context.Context, inputTeam model.InputTeam) (*model.Team, error) {
	panic(fmt.Errorf("not implemented: CreateTeam - createTeam"))
}

// UpdateTeam is the resolver for the updateTeam field.
func (r *mutationResolver) UpdateTeam(ctx context.Context, updateTeam model.UpdateTeam) (*model.Team, error) {
	panic(fmt.Errorf("not implemented: UpdateTeam - updateTeam"))
}

// DeleteTeam is the resolver for the deleteTeam field.
func (r *mutationResolver) DeleteTeam(ctx context.Context, teamID string) (string, error) {
	panic(fmt.Errorf("not implemented: DeleteTeam - deleteTeam"))
}

// Team is the resolver for the team field.
func (r *queryResolver) Team(ctx context.Context, teamID string) ([]*model.Team, error) {
	panic(fmt.Errorf("not implemented: Team - team"))
}
