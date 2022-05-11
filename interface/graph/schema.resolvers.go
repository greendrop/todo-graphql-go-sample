package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	graphgenerated "github.com/greendrop/todo-graphql-go-sample/interface/graph/generated"
	graphmodel "github.com/greendrop/todo-graphql-go-sample/interface/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input graphmodel.NewTodo) (*graphmodel.Todo, error) {
	// panic(fmt.Errorf("not implemented"))
	// todo := &graphmodel.Todo{
	// 	Text: input.Text,
	// 	ID:   fmt.Sprintf("T%d", rand.Int()),
	// 	// User: &graphmodel.User{ID: input.UserID, Name: "user " + input.UserID},
	// 	UserID: input.UserID,
	// }
	// r.todos = append(r.todos, todo)
	// return todo, nil

	todo, _ := r.todoCreateTodoUseCase.Execute(&entity.Todo{
		UserId: 1,
		Title:  &input.Text,
	})

	return &graphmodel.Todo{
		ID:     strconv.FormatInt(todo.Id, 10),
		Text:   *todo.Title,
		Done:   todo.Done,
		UserID: strconv.FormatInt(todo.UserId, 10),
	}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*graphmodel.Todo, error) {
	// panic(fmt.Errorf("not implemented"))
	// return r.todos, nil

	todos, _ := r.todoGetTodoListUseCase.Execute()

	var graphmodelTodos []*graphmodel.Todo
	for _, todo := range *todos {
		graphmodelTodos = append(graphmodelTodos, &graphmodel.Todo{
			ID:     strconv.FormatInt(todo.Id, 10),
			Text:   *todo.Title,
			Done:   todo.Done,
			UserID: strconv.FormatInt(todo.UserId, 10),
		})
	}

	return graphmodelTodos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *graphmodel.Todo) (*graphmodel.User, error) {
	// panic(fmt.Errorf("not implemented"))
	// return &graphmodel.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil

	userId, _ := strconv.ParseInt(obj.UserID, 10, 64)
	user, _ := r.userGetUserUseCase.Execute(userId)

	return &graphmodel.User{
		ID:   strconv.FormatInt(user.Id, 10),
		Name: *user.Name,
	}, nil
}

// Mutation returns graphgenerated.MutationResolver implementation.
func (r *Resolver) Mutation() graphgenerated.MutationResolver { return &mutationResolver{r} }

// Query returns graphgenerated.QueryResolver implementation.
func (r *Resolver) Query() graphgenerated.QueryResolver { return &queryResolver{r} }

// Todo returns graphgenerated.TodoResolver implementation.
func (r *Resolver) Todo() graphgenerated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
