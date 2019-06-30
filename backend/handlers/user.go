package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest/backend/services"
	"net/http"
)

type UserHandler struct {
}

func (h UserHandler) Get(c *gin.Context) {
	us, err := services.GetUsers(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, us)
}


func (h UserHandler) Create(c *gin.Context) {
	u, err := services.CreateUser(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Print(err)
		return
	}
	c.JSON(http.StatusOK, u)
}
