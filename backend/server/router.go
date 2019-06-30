package server

import (
	"gin-sample/backend/database"
	"gin-sample/backend/handlers"
	"gin-sample/backend/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.CORSMiddleware())

	newHealthHandler(router)
	v1 := router.Group("v1")
	{
		newUserHandler(v1)
	}
	return router
}

func newHealthHandler(router *gin.Engine) {
	health := new(handlers.HealthHandler)
	router.GET("health/", health.Status)
}

func newUserHandler(gr *gin.RouterGroup) {
	userGr := gr.Group("users/")
	{
		user := handlers.NewUserHandler(database.GetDB())
		userGr.GET("", user.Get)
		userGr.POST("", user.Create)
		userGr.PUT(":id/", user.Update)
		userGr.DELETE(":id/", user.Delete)
	}
}
