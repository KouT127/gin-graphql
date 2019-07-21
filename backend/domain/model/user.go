// go:generate go run github.com/vektah/dataloaden UserLoader int *github.com/KouT127/gin-sample/backend/domain/model.User
package model

import (
	"github.com/jinzhu/gorm"
)

type UserConnection struct {
	TotalCount int
	PageInfo PageInfo
	Edges    []*UserEdge
}
type UserEdge struct {
	Cursor string
	Node   *User
}
type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	BirthDay string `gorm:"birthday"`
	Gender   string `gorm:"gender"`
	PhotoURL string `gorm:"photo_url"`
	Active   bool   `gorm:"active"`
	Tasks    []Task `gorm:"foreignkey:UserRefer"`
}
