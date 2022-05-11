package persistence

import (
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	"github.com/greendrop/todo-graphql-go-sample/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (p userPersistence) GetById(id int64) (*entity.User, error) {
	var user entity.User
	err := GormDB.Where(&entity.User{Id: id}).Take(&user).Error
	return &user, err
}
