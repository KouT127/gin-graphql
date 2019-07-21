package model

import "github.com/jinzhu/gorm"

type TaskConnection struct {
	TotalCount int
	PageInfo   PageInfo
	Edges      []*TaskEdge
}
type TaskEdge struct {
	Cursor string
	Node   *Task
}
type Task struct {
	gorm.Model
	UserRefer   uint
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
}
