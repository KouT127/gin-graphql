package services

import (
	"github.com/gin-gonic/gin"
	"go-rest/backend/db"
	"go-rest/backend/entities"
)

type User entities.User

func GetUsers(c *gin.Context) ([]User, error) {
	db := db.GetDB()

	var users []User
	db.Find(&users)
	return users, nil
}

func CreateUser(c *gin.Context) (User, error) {
	db := db.GetDB()

	u := User{
		Name:     "",
		BirthDay: "",
		Gender:   "",
		PhotoURL: "",
		Active:   true,
	}
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
