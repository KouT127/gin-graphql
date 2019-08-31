package repository

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
)

func NewItemRepository() *itemRepository {
	return &itemRepository{}
}

type itemRepository struct{}

func (r *itemRepository) FetchItemsCount() (int, error) {
	var cnt int
	db := database.NewDB()
	qs := db.Model(&model.Item{})
	err := qs.Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func (r *itemRepository) FindItems(q *model.Query) (int, []*model.Item, error) {
	var items []*model.Item
	idx, scopes := CalculatePageInfo(q)
	db := database.NewDB()
	rows, err := db.Model(&model.Item{}).Scopes(scopes...).Rows()
	if err != nil {
		return idx, nil, err
	}
	for rows.Next() {
		item := &model.Item{}
		err := db.ScanRows(rows, item)
		if err != nil {
			return idx, nil, err
		}
		items = append(items, item)
	}
	return idx, items, nil
}
