package repository

import (
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
)

type UserRepository interface {
	GetById(id int64) (*entity.User, error)
}
