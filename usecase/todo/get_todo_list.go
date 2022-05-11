package usecase_todo

import (
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	"github.com/greendrop/todo-graphql-go-sample/domain/repository"
)

type TodoGetTodoListUseCase interface {
	Execute() (*[]entity.Todo, error)
}

type todoGetTodoListUseCase struct {
	todoRepository repository.TodoRepository
}

func NewTodoGetTodoListUseCase(todoRepository repository.TodoRepository) TodoGetTodoListUseCase {
	return &todoGetTodoListUseCase{
		todoRepository: todoRepository,
	}
}

func (u todoGetTodoListUseCase) Execute() (*[]entity.Todo, error) {
	return u.todoRepository.GetList()
}
