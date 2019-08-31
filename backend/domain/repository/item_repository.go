package repository

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
)

type ItemRepository interface {
	FetchItemsCount() (int, error)
	FindItems(q *model.Query) (int, []*model.Item, error)
}
