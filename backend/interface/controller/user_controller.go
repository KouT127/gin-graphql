package controller

import (
	"fmt"
	"github.com/KouT127/gin-sample/backend/application/form"
	"github.com/KouT127/gin-sample/backend/application/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

//type UserController interface {
//	Get(c *gin.Context)
//	Create(c *gin.Context)
//}

type UserController struct {
	it interactor.UserInteractor
}

func NewUserController(it interactor.UserInteractor) *UserController {
	return &UserController{
		it: it,
	}
}

func (h UserController) Get(c *gin.Context) {
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

func (h UserController) Create(c *gin.Context) {
	frm := form.UserForm{}
	err := c.Bind(&frm)
	if err != nil {
		c.Status(http.StatusBadRequest)
		print(err.Error())
		return
	}
	res, err := h.it.CreateUser(&frm)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

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
