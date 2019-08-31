package usecase

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	. "github.com/KouT127/gin-sample/backend/interface/graphql/graph"
)

type TaskUsecase interface {
	AllTasks(q *model.Query, id *int) (*TaskConnection, error)
}

type taskUsecase struct{}

func NewTaskUsecase() *taskUsecase {
	return &taskUsecase{}
}

func (tu *taskUsecase) AllTasks(q *model.Query, id *int) (*TaskConnection, error) {
	return &TaskConnection{}, nil
	//var (
	//	cnt   int
	//	edges []*TaskEdge
	//	tasks []*model.Task
	//	err   error
	//)
	//cnt, err = repository.FetchTasksCount()
	//if err != nil {
	//	return nil, err
	//}
	//idx, scopes := repository.CalculatePageInfo(q)
	//tasks, err = repository.FindTasks(scopes)
	//if err != nil {
	//	return nil, err
	//}
	//for _, t := range tasks {
	//	edge := NewTaskEdge(t, idx)
	//	edges = append(edges, edge)
	//	idx++
	//}
	//return NewTaskConnection(cnt, edges), nil
}
