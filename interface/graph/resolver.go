package graph

import (
	graphmodel "github.com/greendrop/todo-graphql-go-sample/interface/graph/model"
	usecase_todo "github.com/greendrop/todo-graphql-go-sample/usecase/todo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos                  []*graphmodel.Todo
	todoGetTodoListUseCase usecase_todo.TodoGetTodoListUseCase
	todoCreateTodoUseCase  usecase_todo.TodoCreateTodoUseCase
}

func NewResolver(todoGetTodoListUseCase usecase_todo.TodoGetTodoListUseCase, todoCreateTodoUseCase usecase_todo.TodoCreateTodoUseCase) *Resolver {
	return &Resolver{
		todoGetTodoListUseCase: todoGetTodoListUseCase,
		todoCreateTodoUseCase:  todoCreateTodoUseCase,
	}
}
