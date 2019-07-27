package graphql

import (
	"context"
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/dataloader"
	. "github.com/jinzhu/gorm"
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

func (r *mutationResolver) AddUser(ctx context.Context, user generated.UserInput) (*generated.AddUserPayload, error) {
	panic("not implemented")
}
func (r *mutationResolver) AddTask(ctx context.Context, input generated.TaskInput) (*generated.AddTaskPayload, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) User(ctx context.Context, id *int) (*graph.User, error) {
	ldr, err := dataloader.CtxLoaders(ctx)
	if err != nil {
		panic(err)
	}
	user, err := ldr.UserById.Load(*id)
	if err != nil {
		panic(err)
	}
	u := graph.NewUser(user)
	return u, nil
}

func (r *queryResolver) Tasks(ctx context.Context, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	db := database.NewDB()
	var cnt int
	var edges []*graph.TaskEdge
	err := db.Model(&model.Task{}).Count(&cnt).Error
	if err != nil {
		panic(err)
	}
	rows, err := db.Model(&model.Task{}).Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		task := &model.Task{}
		err := db.ScanRows(rows, task)
		if err != nil {
			panic(err)
		}
		edge := graph.NewTaskEdge(task)
		edges = append(edges, edge)
	}
	conn := graph.NewTaskConnection(cnt, edges)
	return conn, nil
}

type taskResolver struct{ *Resolver }

func (r *taskResolver) ID(ctx context.Context, obj *graph.Task) (string, error) {
	return obj.ID, nil
}
func (r *taskResolver) User(ctx context.Context, obj *graph.Task) (*graph.User, error) {
	id := int(obj.UserRefer)
	return r.Query().User(ctx, &id)
}

type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *graph.User) (string, error) {
	return obj.ID, nil
}
func (r *userResolver) Tasks(ctx context.Context, obj *graph.User, first *int, after *string, last *int, before *string, query *string) (*graph.TaskConnection, error) {
	var scopes []func(*DB) *DB
	var cnt int
	var edges []*graph.TaskEdge
	db := database.NewDB()
	
	q, err := dataloader.NewQuery(first, after, last, before, query)
	if err != nil {
		panic(err)
	}
	qs := db.Model(&model.Task{})
	err = qs.Count(&cnt).Error
	if err != nil {
		panic(err)
	}
	if q.First != 0 {
		scopes = append(scopes, ByCount(q.First))
	}
	if q.Last != 0 {
		scopes = append(scopes, ByCount(q.Last))
	}
	if q.After != 0 {
		scopes = append(scopes, ByOffset(q.After))
	}
	if q.Before != 0 {
		scopes = append(scopes, ByOffset(q.Before))
	}
	if q.Keyword != "" {
		scopes = append(scopes, ByKeyword(q.Keyword))
	}
	rows, err := db.Model(&model.Task{}).Scopes(scopes...).Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		task := &model.Task{}
		err := db.ScanRows(rows, task)
		if err != nil {
			panic(err)
		}
		edge := graph.NewTaskEdge(task)
		edges = append(edges, edge)
	}
	conn := graph.NewTaskConnection(cnt, edges)
	return conn, nil
}

func ByKeyword(keyword string) func(db *DB) *DB {
	return func(db *DB) *DB {
		return db.Where("keyword = ?", keyword)
	}
}

func ByCount(count int) func(db *DB) *DB {
	return func(db *DB) *DB {
		return db.Limit(count)
	}
}

func ByOffset(offset int) func(db *DB) *DB {
	return func(db *DB) *DB {
		return db.Offset(offset)
	}
}
