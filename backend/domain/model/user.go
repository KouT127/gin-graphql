// go:generate go run github.com/vektah/dataloaden UserLoader int *github.com/KouT127/gin-sample/backend/domain/model.User
package user

import (
	"github.com/KouT127/gin-sample/backend/domain/model/task"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string      `gorm:"name"`
	BirthDay string      `gorm:"birthday"`
	Gender   string      `gorm:"gender"`
	PhotoURL string      `gorm:"photo_url"`
	Active   bool        `gorm:"active"`
	Tasks    []task.Task `gorm:"foreignkey:UserRefer"`
}

func (User) IsNode() {}
