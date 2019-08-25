package resolver

import (
	"context"
	"github.com/KouT127/gin-sample/backend/application/usecase"
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/interface/controller"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/dataloader"
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

func (r *Resolver) Item() generated.ItemResolver {
	return &itemResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddUser(ctx context.Context, user generated.UserInput) (*generated.AddUserPayload, error) {
	uc := controller.NewUserController()
	payload, err := uc.AddTask(&user)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func (r *mutationResolver) AddTask(ctx context.Context, input generated.TaskInput) (*generated.AddTaskPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Items(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.ItemConnection, error) {
	q, err := model.NewQuery(first, after, last, before, query)
	if err != nil {
		return nil, err
	}
	uc := usecase.NewItemUsecase()
	conn, err := uc.AllItems(q, nil)
	if err != nil {
		return conn, err
	}
	return conn, nil
}

func (r *queryResolver) User(ctx context.Context, id *int) (*graph.User, error) {
	ldr, err := dataloader.CtxLoaders(ctx)
	if err != nil {
		return &graph.User{}, err
	}
	user, err := ldr.UserById.Load(*id)
	if err != nil {
		return &graph.User{}, err
	}
	u := graph.NewUser(user)
	return u, nil
}

func (r *queryResolver) Tasks(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	q, err := model.NewQuery(first, after, last, before, query)
	if err != nil {
		return &graph.TaskConnection{}, err
	}
	tu := usecase.NewTaskUsecase()
	conn, err := tu.AllTasks(q, nil)
	if err != nil {
		return conn, err
	}
	return conn, nil
}
