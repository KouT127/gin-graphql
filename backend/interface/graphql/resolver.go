// go:generate gqlgen
package graphql

import (
	"context"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"

	"github.com/KouT127/gin-sample/backend/domain/model"
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

func (r *mutationResolver) CreateUser(ctx context.Context, user graph.UserInput) (*model.User, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateTask(ctx context.Context, input graph.TaskInput) (*model.Task, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Task(ctx context.Context, id *string) (*model.Task, error) {
	panic("not implemented")
}
func (r *queryResolver) Tasks(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	db := database.NewDB()
	var task *model.Task
	rows, err := db.Model(&model.Task{}).Rows()
	if err != nil {
		panic(err)
	}
	var list []*model.User
	for rows.Next() {
		mem := &model.User{}
		err := db.ScanRows(rows, &mem)
		if err != nil {
			panic(err)
		}
		list = append(list, mem)
	}
	edge := &graph.TaskEdge{
		Cursor: "",
		Node:   task,
	}
	edges := []*graph.TaskEdge{}
	edges = append(edges, edge)
	endcur := "test"

	pg := &graph.PageInfo{
		EndCursor:   &endcur,
		HasNextPage: true,
	}
	con := &graph.TaskConnection{
		TotalCount: 0,
		Edges:      edges,
		PageInfo:   pg,
	}
	return con, nil
}
func (r *queryResolver) Users(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.UserConnection, error) {
	panic("not implemented")
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) ID(ctx context.Context, obj *model.Task) (string, error) {
	panic("not implemented")
}

func (r *taskResolver) User(ctx context.Context, obj *model.Task) (*model.User, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic("not implemented")
}
