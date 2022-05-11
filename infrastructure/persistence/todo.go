package persistence

import (
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	"github.com/greendrop/todo-graphql-go-sample/domain/repository"
)

type todoPersistence struct{}

func NewTodoPersistence() repository.TodoRepository {
	return &todoPersistence{}
}

func (p todoPersistence) GetList() (*[]entity.Todo, error) {
	var todos []entity.Todo
	err := GormDB.Order("id asc").Find(&todos).Error
	return &todos, err
}

func (p todoPersistence) GetById(id int64) (*entity.Todo, error) {
	var todo entity.Todo
	err := GormDB.Where(&entity.Todo{Id: id}).Take(&todo).Error
	return &todo, err
}

func (p todoPersistence) Create(todo *entity.Todo) (*entity.Todo, error) {
	err := GormDB.Create(todo).Error
	return todo, err
}

func (p todoPersistence) Update(todo *entity.Todo) (*entity.Todo, error) {
	err := GormDB.Model(todo).Updates(entity.Todo{Title: todo.Title, Done: todo.Done}).Error
	return todo, err
}

func (p todoPersistence) Delete(todo *entity.Todo) (*entity.Todo, error) {
	err := GormDB.Delete(todo).Error
	return todo, err
}
