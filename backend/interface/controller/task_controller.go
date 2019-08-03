package controller

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	. "github.com/KouT127/gin-sample/backend/interface/graphql/graph"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/dataloader"
	. "github.com/jinzhu/gorm"
)

type TaskController interface {
	AllTasks(q *dataloader.Query) (*TaskConnection, error)
}

type taskController struct{}

func NewTaskController() *taskController {
	return &taskController{}
}

func (taskController) AllTasks(q *dataloader.Query) (*TaskConnection, error) {
	var scopes []func(db *DB) *DB
	var cnt, idx int
	var edges []*TaskEdge
	var err error
	db := database.NewDB()
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
		scopes = append(scopes, OrderBySort("id DESC"))
	}
	if q.After != 0 {
		idx += q.After
	}
	if q.Before != 0 {
		idx += q.Before
	}
	if q.Keyword != "" {
		scopes = append(scopes, ByKeyword(q.Keyword))
	}
	scopes = append(scopes, ByOffset(idx))
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
		edge := NewTaskEdge(task, idx)
		edges = append(edges, edge)
		idx++
	}
	return NewTaskConnection(cnt, edges), nil
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

func OrderBySort(sort string) func(db *DB) *DB {
	return func(db *DB) *DB {
		return db.Order(sort)
	}
}
