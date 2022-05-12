package loader

import (
	"context"
	"fmt"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	"github.com/greendrop/todo-graphql-go-sample/infrastructure/persistence"
)

func UsersLoader(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	userIDs := make([]int64, len(keys))
	for ix, key := range keys {
		userId, _ := strconv.ParseInt(key.String(), 10, 64)
		userIDs[ix] = userId
	}

	var users []entity.User
	persistence.GormDB.Find(&users, userIDs)

	userById := map[int64]*entity.User{}
	for _, user := range users {
		userById[user.Id] = &user
	}

	output := make([]*dataloader.Result, len(keys))
	for index, userKey := range keys {
		userId, _ := strconv.ParseInt(userKey.String(), 10, 64)
		user, ok := userById[userId]
		if ok {
			output[index] = &dataloader.Result{Data: user, Error: nil}
		} else {
			err := fmt.Errorf("user not found %s", userKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}
