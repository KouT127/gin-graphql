package resolver

import (
	"context"
	"github.com/KouT127/gin-sample/backend/application/usecase"
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
	"strconv"
)

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *graph.User) (string, error) {
	return obj.ID, nil
}
func (r *userResolver) Tasks(ctx context.Context, obj *graph.User, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	q, err := model.NewQuery(first, after, last, before, query)
	if err != nil {
		return nil, err
	}
	uid, err := strconv.Atoi(obj.ID)
	if err != nil {
		return nil, err
	}
	tu := usecase.NewTaskUsecase()
	conn, err := tu.AllTasks(q, &uid)
	if err != nil {
		return nil, err
	}
	return conn, nil
}