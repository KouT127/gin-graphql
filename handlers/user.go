package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-rest/services"
	"net/http"
)

type UserHandler struct {
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
