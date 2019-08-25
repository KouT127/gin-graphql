package usecase

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/repository"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
)

type ItemUsecase interface {
	AllItems(q *model.Query, id *int) (*graph.ItemConnection, error)
}

type itemUsecase struct{}

func NewItemUsecase() *itemUsecase {
	return &itemUsecase{}
}

func (iu *itemUsecase) AllItems(q *model.Query, id *int) (*graph.ItemConnection, error) {
	var (
		cnt   int
		edges []*graph.ItemEdge
		items []*model.Item
		err   error
	)
	cnt, err = repository.FetchItemsCount()
	if err != nil {
		return nil, err
	}
	idx, scopes := repository.CalculatePageInfo(q)
	items, err = repository.FindItems(scopes)
	if err != nil {
		return nil, err
	}
	for _, i := range items {
		edge := graph.NewItemEdge(i, idx)
		edges = append(edges, edge)
		idx++
	}
	return graph.NewItemConnection(cnt, edges), nil
}
