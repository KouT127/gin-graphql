package form

import (
	"github.com/jinzhu/gorm"
)

type Paginator interface {
	Paging()
}

type Pagination struct {
	Page  int `form:"page" binding:`
	Limit int `form:"limit" binding:`
}

func (pg Pagination) Paging(db *gorm.DB) *gorm.DB {
	l := 10
	p := 0

	if pg.Page != 0 {
		p = pg.Page -1
	}
	if pg.Limit != 0 {
		l = pg.Limit
	}
	c := l*p
	return db.Offset(c).Limit(l)
}
