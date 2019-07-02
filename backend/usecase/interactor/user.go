//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package interactor

import (
	"gin-sample/backend/domain/model"
	"gin-sample/backend/interface/gateway"
)

type User model.User

type UserInteractor interface {
	GetUsers() ([]User, error)
	CreateUser() (User, error)
}

type userInteractor struct {
	ur *gateway.UserRepository
}

func NewUserInteractor(ur gateway.UserRepository) *userInteractor {
	return &userInteractor{
		ur: &ur,
	}
}

func (i *userInteractor) GetUsers() ([]User, error) {
	var users []User
	//i.ur.FindAll(&users)
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
	//if err := s.db.Create(&u).Error; err != nil {
	//	return u, err
	//}
	return u, nil
}

//
//func (s *userInteractor) UpdateUser(id string) (User, error) {
//	var u User
//	s.db.First(&u, "id = ?", id)
//	u.Name = "updated"
//
//	if err := s.db.Save(&u).Error; err != nil {
//		return u, err
//	}
//	return u, nil
//}
//
//func (s *userInteractor) DeleteUser(id string) (User, error) {
//	var u User
//	s.db.First(&u, "id = ?", id)
//	if err := s.db.Delete(&u).Error; err != nil {
//		return u, err
//	}
//	return u, nil
//}
