package repository

import (
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
)

type TodoRepository interface {
	GetList() (*[]entity.Todo, error)
	GetById(id int64) (*entity.Todo, error)
	Create(task *entity.Todo) (*entity.Todo, error)
	Update(task *entity.Todo) (*entity.Todo, error)
	Delete(task *entity.Todo) (*entity.Todo, error)
}
