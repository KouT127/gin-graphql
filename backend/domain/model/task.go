package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	UserRefer   uint
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
}
