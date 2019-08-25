package repository

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	. "github.com/jinzhu/gorm"
)

func FetchItemsCount() (int, error) {
	var cnt int
	db := database.NewDB()
	qs := db.Model(&model.Item{})
	err := qs.Count(&cnt).Error
	if err != nil {
		return 0, err
	}
	return cnt, nil
}

func FindItems(scopes []func(db *DB) *DB) ([]*model.Item, error) {
	var items []*model.Item
	db := database.NewDB()
	rows, err := db.Model(&model.Item{}).Scopes(scopes...).Rows()
	if err != nil {
		return []*model.Item{}, err
	}
	for rows.Next() {
		item := &model.Item{}
		err := db.ScanRows(rows, item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
