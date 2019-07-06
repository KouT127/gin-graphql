package presenter

import (
	"gin-sample/backend/domain/model"
	"gin-sample/backend/usecase/response"
)

type UserPresenter interface {
	PresentUsers(us []*model.User) response.UsersResponse
}

type userPresenter struct{}

func NewUserPresenter() *userPresenter {
	return &userPresenter{}
}

func (h userPresenter) PresentUsers(us []*model.User) response.UsersResponse {
	var usrAry []*response.UserResponse
	for _, u := range us {
		user := response.UserResponse{
			ID:     u.ID,
			Gender: u.Gender,
			Name:   u.Name,
		}
		usrAry = append(usrAry, &user)
	}
	res := response.UsersResponse{Users: usrAry}
	return res
}

func (h userPresenter) PresentUser(u *model.User) response.UserResponse {

	res := response.UserResponse{
		ID:     u.ID,
		Name:   u.Name,
		Gender: u.Gender,
	}
	return res
}

func (h userPresenter) PresentError(err error) {
}
