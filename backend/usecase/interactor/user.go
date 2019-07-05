//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package interactor

import (
	"fmt"
	"gin-sample/backend/domain/model"
	"gin-sample/backend/interface/gateway"
)

type UserInteractor interface {
	GetUsers() ([]*model.User, error)
	CreateUser() (model.User, error)
}

type userInteractor struct {
	ur gateway.UserRepository
}

func NewUserInteractor(ur gateway.UserRepository) *userInteractor {
	return &userInteractor{
		ur: ur,
	}
}

func (i *userInteractor) GetUsers() ([]*model.User, error) {
	users, err := i.ur.FindAll()
	if err != nil {
		fmt.Print("Interactor")
	}
	return users, nil
}

func (s *userInteractor) CreateUser() (model.User, error) {

	u := model.User{
		Name:     "",
		BirthDay: "",
		Gender:   "",
		PhotoURL: "",
		Active:   true,
	}
	if _, err := s.ur.Create(&u); err != nil {
		fmt.Print("Interactor")
	}
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
