package repository

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	. "github.com/jinzhu/gorm"
)

type ItemRepository interface {
	FetchItemsCount() (int, error)
	FindItems(scopes []func(db *DB) *DB) ([]*model.Item, error)
}
