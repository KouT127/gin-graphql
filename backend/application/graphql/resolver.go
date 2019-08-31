package graphql

import (
	"context"
	generated2 "github.com/KouT127/gin-sample/backend/application/graphql/generated"
	graph2 "github.com/KouT127/gin-sample/backend/application/graphql/graph"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Item() generated2.ItemResolver {
	return &itemResolver{r}
}
func (r *Resolver) Mutation() generated2.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated2.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Task() generated2.TaskResolver {
	return &taskResolver{r}
}
func (r *Resolver) User() generated2.UserResolver {
	return &userResolver{r}
}

type itemResolver struct{ *Resolver }

func (r *itemResolver) Price(ctx context.Context, obj *graph2.Item) (float64, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUser(ctx context.Context, user generated2.UserInput) (*generated2.AddUserPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddTask(ctx context.Context, input generated2.TaskInput) (*generated2.AddTaskPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id *int) (*graph2.User, error) {
	panic("not implemented")
}
func (r *queryResolver) Tasks(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph2.TaskConnection, error) {
	panic("not implemented")
}
func (r *queryResolver) Items(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph2.ItemConnection, error) {
	panic("not implemented")
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) User(ctx context.Context, obj *graph2.Task) (*graph2.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Tasks(ctx context.Context, obj *graph2.User, first *int, after *string, last *int, before *string, query *string) (*graph2.TaskConnection, error) {
	panic("not implemented")
}
func (r *userResolver) Cart(ctx context.Context, obj *graph2.User) (*graph2.Cart, error) {
	panic("not implemented")
}
