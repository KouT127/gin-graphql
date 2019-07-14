//go:generate go run github.com/99designs/gqlgen
package graphql

import (
	"context"
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/graphql/generated"
	"github.com/KouT127/gin-sample/backend/infrastracture/graphql/graph"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Task() generated.TaskResolver {
	return &taskResolver{r}
}
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateUser(ctx context.Context, user graph.NewUser) (*model.User, error) {
	panic("implement me")
}

func (r *mutationResolver) CreateTask(ctx context.Context, input graph.NewTask) (*model.Task, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) FindTaskByID(ctx context.Context, id *int) (*model.Task, error) {
	panic("implement me")
}

func (r *queryResolver) FindTasks(ctx context.Context) ([]*model.Task, error) {
	panic("implement me")
}

func (r *queryResolver) FindUsers(ctx context.Context) ([]*model.User, error) {
	panic("implement me")
}

func (r *queryResolver) GetTaskByID(ctx context.Context, id *int) (*model.Task, error) {
	return &model.Task{
		UserRefer:   1,
		Title:       "1",
		Description: "1",
	},nil
}
func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	panic("not implemented")
}
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic("not implemented")
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) ID(ctx context.Context, obj *model.Task) (int, error) {
	panic("not implemented")
}
func (r *taskResolver) User(ctx context.Context, obj *model.Task) (*model.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *model.User) (int, error) {
	panic("not implemented")
}
