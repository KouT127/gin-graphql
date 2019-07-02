package gateway

import (
	"gin-sample/backend/domain/model"
	"github.com/jinzhu/gorm"
)

type User model.User

type UserRepository interface {
	FindAll() ([]User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}
func (ur *userRepository) FindAll() ([]User, error) {
	var users []User
	ur.db.Find(&users)
	return users, nil
}
