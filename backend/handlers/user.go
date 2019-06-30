package handlers

import (
	"fmt"
	"gin-sample/backend/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type UserHandler interface{}

type userHandler struct {
	srv services.UserService
}

func NewUserHandler(db *gorm.DB) *userHandler {
	return &userHandler{
		srv: services.NewUserService(db),
	}
}

func (h userHandler) Get(c *gin.Context) {
	us, err := h.srv.GetUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, us)
}

func (h userHandler) Create(c *gin.Context) {
	u, err := h.srv.GetUsers()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h userHandler) Update(c *gin.Context) {
	id := c.Param("id")
	u, err := h.srv.UpdateUser(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func (h userHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := h.srv.DeleteUser(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}
