//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package services

import (
	"gin-sample/backend/domain/entities"
	"github.com/jinzhu/gorm"
)

type User entities.User

type UserInteractor interface {
	GetUsers() ([]User, error)
	CreateUser() (User, error)
	UpdateUser(id string) (User, error)
	DeleteUser(id string) (User, error)
}

type userInteractor struct {
	db *gorm.DB
}

func NewUserInteractor(db *gorm.DB) *userInteractor {
	return &userInteractor{
		db: db,
	}
}

func (s *userInteractor) GetUsers() ([]User, error) {
	var users []User
	s.db.Find(&users)
	return users, nil
}

func (s *userInteractor) CreateUser() (User, error) {
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

func (s *userInteractor) UpdateUser(id string) (User, error) {
	var u User
	s.db.First(&u, "id = ?", id)
	u.Name = "updated"

	if err := s.db.Save(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

func (s *userInteractor) DeleteUser(id string) (User, error) {
	var u User
	s.db.First(&u, "id = ?", id)
	if err := s.db.Delete(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}
