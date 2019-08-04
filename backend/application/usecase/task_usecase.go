package usecase

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/repository"
	. "github.com/KouT127/gin-sample/backend/interface/graphql/graph"
	"github.com/jinzhu/gorm"
)

type TaskUsecase interface {
	AllTasks(q *model.Query, id *int) (*TaskConnection, error)
}

type taskUsecase struct{}

func NewTaskUsecase() *taskUsecase {
	return &taskUsecase{}
}

func calculatePageInfo(q *model.Query) (int, []func(db *gorm.DB) *gorm.DB) {
	var (
		scopes []func(db *gorm.DB) *gorm.DB
		idx    int
	)
	if q.First != 0 {
		scopes = append(scopes, repository.ByCount(q.First))
	}
	if q.Last != 0 {
		scopes = append(scopes, repository.ByCount(q.Last))
		scopes = append(scopes, repository.OrderBySort("id DESC"))
	}
	if q.After != 0 {
		idx += q.After
	}
	if q.Before != 0 {
		idx += q.Before
	}
	if q.Keyword != "" {
		scopes = append(scopes, repository.ByKeyword(q.Keyword))
	}
	scopes = append(scopes, repository.ByOffset(idx))
	return idx, scopes
}

func (tu *taskUsecase) AllTasks(q *model.Query, id *int) (*TaskConnection, error) {
	var (
		cnt   int
		edges []*TaskEdge
		tasks  []*model.Task
		err   error
	)
	cnt, err = repository.FetchTasksCount()
	if err != nil {
		return nil, err
	}
	idx, scopes := calculatePageInfo(q)
	tasks, err = repository.FindTasks(scopes)
	if err != nil {
		return nil, err
	}
	for _, t := range tasks {
		edge := NewTaskEdge(t, idx)
		edges = append(edges, edge)
		idx++
	}
	return NewTaskConnection(cnt, edges), nil
}

