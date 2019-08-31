package usecase

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/domain/repository"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
)

type ItemUsecase interface {
	AllItems(q *model.Query, id *int) (*graph.ItemConnection, error)
}

type itemUsecase struct {
	ir repository.ItemRepository
}

func NewItemUsecase(ir repository.ItemRepository) *itemUsecase {
	return &itemUsecase{
		ir: ir,
	}
}

func (iu *itemUsecase) AllItems(q *model.Query, id *int) (*graph.ItemConnection, error) {
	var (
		cnt, idx int
		edges    []*graph.ItemEdge
		items    []*model.Item
		err      error
	)
	cnt, err = iu.ir.FetchItemsCount()
	if err != nil {
		return nil, err
	}
	idx, items, err = iu.ir.FindItems(q)
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
