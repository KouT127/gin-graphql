package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go-rest/db"
	"go-rest/entities"
	"time"
)

type User entities.User

func CreateUser(c *gin.Context) (User, error) {
	db := db.GetDB()

	u := User{
		Model:     gorm.Model{
			ID:        2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		Name:      "",
		BirthDay:  "",
		Gender:    "",
		PhotoURL:  "",
		Time:      0,
		Active:    true,
	}
	if err := db.Create(u).Error; err != nil {
		return u, err
	}
	return u, nil
}
