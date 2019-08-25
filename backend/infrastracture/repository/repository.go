package repository

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	. "github.com/jinzhu/gorm"
)

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

func CalculatePageInfo(q *model.Query) (int, []func(db *DB) *DB) {
	var (
		scopes []func(db *DB) *DB
		idx    int
	)
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
	return idx, scopes
}
