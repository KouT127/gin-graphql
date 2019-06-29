package server

import (
	"github.com/gin-gonic/gin"
	"go-rest/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(handlers.HealthHandler)
	router.GET("/health", health.Status)
	user := new(handlers.UserHandler)
	router.GET("/user", user.Create)
	//router.Use(middlewares.AuthMiddleware())
	//
	//v1 := router.Group("v1")
	//{
	//	userGroup := v1.Group("user")
	//	{
	//		user := new(controllers.UserController)
	//		userGroup.GET("/:id", user.Retrieve)
	//	}
	//}
	return router
}
