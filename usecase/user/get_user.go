package usecase_user

import (
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	"github.com/greendrop/todo-graphql-go-sample/domain/repository"
)

type UserGetUserUseCase interface {
	Execute(id int64) (*entity.User, error)
}

type userGetUserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserGetUserUseCase(userRepository repository.UserRepository) UserGetUserUseCase {
	return &userGetUserUseCase{
		userRepository: userRepository,
	}
}

func (u userGetUserUseCase) Execute(id int64) (*entity.User, error) {
	return u.userRepository.GetById(id)
}
