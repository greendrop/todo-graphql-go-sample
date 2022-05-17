package graph2

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	graphgenerated2 "github.com/greendrop/todo-graphql-go-sample/interface/graph2/generated"
	graphmodel2 "github.com/greendrop/todo-graphql-go-sample/interface/graph2/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input graphmodel2.NewTodo) (*graphmodel2.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*graphmodel2.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns graphgenerated2.MutationResolver implementation.
func (r *Resolver) Mutation() graphgenerated2.MutationResolver { return &mutationResolver{r} }

// Query returns graphgenerated2.QueryResolver implementation.
func (r *Resolver) Query() graphgenerated2.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
