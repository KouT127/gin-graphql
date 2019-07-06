package controller

import (
	"fmt"
	"gin-sample/backend/usecase/form"
	"gin-sample/backend/usecase/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	Get(c *gin.Context)
	Create(c *gin.Context)
}

type userController struct {
	it interactor.UserInteractor
}

func NewUserController(it interactor.UserInteractor) *userController {
	return &userController{
		it: it,
	}
}

func (h userController) Get(c *gin.Context) {
	pf := form.Pagination{}
	err := c.Bind(&pf)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	us, err := h.it.GetUsers(&pf)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, us)
}

func (h userController) Create(c *gin.Context) {
	frm := form.UserForm{}
	err := c.Bind(&frm)
	if err != nil {
		c.Status(http.StatusBadRequest)
		print(err.Error())
		return
	}
	u, err := h.it.CreateUser(&frm)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, u)
}

//func (h userController) Update(c *gin.Context) {
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
//func (h userController) Delete(c *gin.Context) {
//	id := c.Param("id")
//	_, err := h.it.DeleteUser(id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusInternalServerError)
//		fmt.Print(err)
//		return
//	}
//	c.JSON(http.StatusOK, "ok")
//}
