//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=$GOPACKAGE
package interactor

import (
	"fmt"
	"github.com/KouT127/gin-sample/backend/application/form"
	"github.com/KouT127/gin-sample/backend/application/response"
	"github.com/KouT127/gin-sample/backend/domain/repository"
	"github.com/KouT127/gin-sample/backend/interface/presenter"
)

type UserInteractor interface {
	GetUsers(pf *form.Pagination) (response.UsersResponse, error)
	CreateUser(frm *form.UserForm) (response.UserResponse, error)
}

type userInteractor struct {
	ur repository.UserRepository
	up presenter.UserPresenter
}

func NewUserInteractor(ur repository.UserRepository, up presenter.UserPresenter) *userInteractor {
	return &userInteractor{
		ur: ur,
		up: up,
	}
}

func (i *userInteractor) GetUsers(pf *form.Pagination) (response.UsersResponse, error) {
	users, err := i.ur.FindAll(pf)
	if err != nil {
		fmt.Print("Interactor")
	}
	mp := i.ur.GetUserMaxPage(pf.Limit)
	res := i.up.PresentUsers(users, mp)
	return res, nil
}

func (i *userInteractor) CreateUser(frm *form.UserForm) (response.UserResponse, error) {
	u, err := i.ur.Create(frm)
	if err != nil {
		fmt.Print("Interactor")
	}
	res := i.up.PresentUser(u)
	return res, nil
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
