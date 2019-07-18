package presenter

import (
	"github.com/KouT127/gin-sample/backend/application/response"
	"github.com/KouT127/gin-sample/backend/domain/model/user"
)

type UserPresenter interface {
	PresentUsers(us []*user.User, mp int) response.UsersResponse
	PresentUser(u *user.User) response.UserResponse
}

type userPresenter struct{}

func NewUserPresenter() *userPresenter {
	return &userPresenter{}
}

func (h userPresenter) PresentUsers(us []*user.User, maxPage int) response.UsersResponse {
	var usrAry []*response.UserResponse
	for _, u := range us {
		user := response.UserResponse{
			ID:     u.ID,
			Gender: u.Gender,
			Name:   u.Name,
		}
		usrAry = append(usrAry, &user)
	}
	res := response.UsersResponse{MaxPage: maxPage, Users: usrAry}
	return res
}

func (h userPresenter) PresentUser(u *user.User) response.UserResponse {
	res := response.UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Gender: u.Gender,
	}
	return res
}

func (h userPresenter) PresentError(err error) {
}
