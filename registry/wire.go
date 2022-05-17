//go:build wireinject
// +build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/greendrop/todo-graphql-go-sample/infrastructure/persistence"
	"github.com/greendrop/todo-graphql-go-sample/interface/graph"
	usecase_todo "github.com/greendrop/todo-graphql-go-sample/usecase/todo"
	usecase_user "github.com/greendrop/todo-graphql-go-sample/usecase/user"
)

func InitializeGraphResolver() *graph.Resolver {
	wire.Build(
		persistence.NewTodoPersistence,
		persistence.NewUserPersistence,
		usecase_todo.NewTodoGetTodoListUseCase,
		usecase_todo.NewTodoCreateTodoUseCase,
		usecase_user.NewUserGetUserUseCase,
		graph.NewResolver,
	)

	return &graph.Resolver{}
}
