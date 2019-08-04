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
		Name:     user.Name,
		Gender:   user.Gender,
		Active:   true,
	}
	db.Save(&u)
	for _, task := range user.Tasks {
		t := model.Task{
			UserRefer:   u.ID,
			Title:       task.Title,
			Description: task.Description,
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

//func (h UserController) Get(c *gin.Context) {
//	//pf := form.Pagination{}
//	//err := c.Bind(&pf)
//	//if err != nil {
//	//	fmt.Printf(err.Error())
//	//	return
//	//}
//	//us, err := h.it.GetUsers(&pf)
//	//if err != nil {
//	//	c.AbortWithStatus(http.StatusInternalServerError)
//	//	fmt.Print(err)
//	//	return
//	//}
//	u := model.User{
//		Name:     "c",
//		BirthDay: "d",
//		Gender:   "f",
//		PhotoURL: "e",
//		Active:   false,
//	}
//	c.JSON(http.StatusOK, u)
//}
//
//func (h UserController) Create(c *gin.Context) {
//	frm := form.UserForm{}
//	err := c.Bind(&frm)
//	if err != nil {
//		c.Status(http.StatusBadRequest)
//		print(err.Error())
//		return
//	}
//	res, err := h.it.CreateUser(&frm)
//	if err != nil {
//		c.AbortWithStatus(http.StatusInternalServerError)
//		fmt.Print(err)
//		return
//	}
//	c.JSON(http.StatusOK, res)
//}

//func (h UserController) Update(c *gin.Context) {
//	id := c.Param("id")
//	u, err := h.it.UpdateUser(id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusInternalServerError)
//		fmt.Print(err)
//		return
//	}
//	c.JSON(http.StatusOK, u)
//}
//
//func (h UserController) Delete(c *gin.Context) {
//	id := c.Param("id")
//	_, err := h.it.DeleteUser(id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusInternalServerError)
//		fmt.Print(err)
//		return
//	}
//	c.JSON(http.StatusOK, "ok")
//}
