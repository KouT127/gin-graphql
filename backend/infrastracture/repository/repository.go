package repository

import (
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
