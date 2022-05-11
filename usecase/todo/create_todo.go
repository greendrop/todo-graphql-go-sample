package usecase_todo

import (
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	"github.com/greendrop/todo-graphql-go-sample/domain/repository"
)

type TodoCreateTodoUseCase interface {
	Execute(todo *entity.Todo) (*entity.Todo, error)
}

type todoCreateTodoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewTodoCreateTodoUseCase(todoRepository repository.TodoRepository) TodoCreateTodoUseCase {
	return &todoCreateTodoUseCase{
		todoRepository: todoRepository,
	}
}

func (u todoCreateTodoUseCase) Execute(todo *entity.Todo) (*entity.Todo, error) {
	return u.todoRepository.Create(todo)
}
