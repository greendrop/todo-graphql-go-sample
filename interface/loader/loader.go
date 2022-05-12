package loader

import (
	"context"
	"net/http"
	"strconv"

	"github.com/graph-gophers/dataloader"
	"github.com/greendrop/todo-graphql-go-sample/domain/entity"
	graphmodel "github.com/greendrop/todo-graphql-go-sample/interface/graph/model"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	UserLoader *dataloader.Loader
}

func NewLoaders() *Loaders {
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(UsersLoader),
	}
	return loaders
}

func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func GetUser(ctx context.Context, userID string) (*graphmodel.User, error) {
	loaders := For(ctx)
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}

	user := result.(*entity.User)

	return &graphmodel.User{
		ID:   strconv.FormatInt(user.Id, 10),
		Name: *user.Name,
	}, nil
}
