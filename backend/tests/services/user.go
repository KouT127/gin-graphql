package services

import (
	"gin-sample/backend/entities"
	services "gin-sample/backend/services/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserService(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := services.NewMockUserSrv(ctrl)

	u := entities.User{
		Name:     "",
		BirthDay: "",
		Gender:   "",
		PhotoURL: "",
		Active:   true,
	}
	m.
		EXPECT().
		GetUsers(gomock.Eq("c")).
		Return(u)
}
