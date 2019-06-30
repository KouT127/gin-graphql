//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package services

import (
	"gin-sample/backend/entities"
	"github.com/jinzhu/gorm"
)

type User entities.User

type UserService interface {
	GetUsers() ([]User, error)
	CreateUser() (User, error)
	UpdateUser(id string) (User, error)
	DeleteUser(id string) (User, error)
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{
		db: db,
	}
}

func (s *userService) GetUsers() ([]User, error) {
	var users []User
	s.db.Find(&users)
	return users, nil
}

func (s *userService) CreateUser() (User, error) {
	u := User{
		Name:     "",
		BirthDay: "",
		Gender:   "",
		PhotoURL: "",
		Active:   true,
	}
	if err := s.db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (s *userService) UpdateUser(id string) (User, error) {
	var u User
	s.db.First(&u, "id = ?", id)
	u.Name = "updated"

	if err := s.db.Save(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (s *userService) DeleteUser(id string) (User, error) {
	var u User
	s.db.First(&u, "id = ?", id)
	if err := s.db.Delete(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
