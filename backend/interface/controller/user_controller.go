package controller

import (
	"github.com/KouT127/gin-sample/backend/domain/model"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/graphql/graph"
	"github.com/KouT127/gin-sample/backend/util"
	"strconv"
)

type UserController interface {
	AddTask(user *generated.UserInput) (*generated.AddUserPayload, error)
}

type userController struct{}

func NewUserController() *userController {
	return &userController{}
}

func (uc *userController) AddTask(user *generated.UserInput) (*generated.AddUserPayload, error) {
	db := database.NewDB()
	u := model.User{
		Name:   user.Name,
		Gender: user.Gender,
		Active: true,
	}
	db.Save(&u)
	for _, task := range user.Tasks {
		t := model.Task{
			UserRefer:   u.ID,
			Title:       task.Title,
			Description: task.Description,
			DeletedAt:   nil,
		}
		db.Save(&t)
	}
	id := strconv.Itoa(int(u.ID))
	ecd := util.Base64Encode("user:" + id)
	usr := graph.User{
		ID:       id,
		Name:     u.Name,
		BirthDay: u.Gender,
		Active:   true,
	}
	payload := generated.AddUserPayload{
		ClientMutationID: &ecd,
		User:             &usr,
	}
	return &payload, nil
}
