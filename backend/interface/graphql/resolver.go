package graphql

import (
	"context"

	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Item() generated.ItemResolver {
	return &itemResolver{r}
}
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

type itemResolver struct{ *Resolver }

func (r *itemResolver) Price(ctx context.Context, obj *graph.Item) (float64, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUser(ctx context.Context, user generated.UserInput) (*generated.AddUserPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddTask(ctx context.Context, input generated.TaskInput) (*generated.AddTaskPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id *int) (*graph.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Tasks(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	panic("not implemented")
}
func (r *queryResolver) Items(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.ItemConnection, error) {
	panic("not implemented")
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) User(ctx context.Context, obj *graph.Task) (*graph.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Tasks(ctx context.Context, obj *graph.User, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	panic("not implemented")
}
func (r *userResolver) Cart(ctx context.Context, obj *graph.User) (*graph.Cart, error) {
	panic("not implemented")
}
