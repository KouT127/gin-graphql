package server

import (
	"gin-sample/backend/infrastracture/database"
	"gin-sample/backend/infrastracture/handlers"
	"gin-sample/backend/infrastracture/middlewares"
	"gin-sample/backend/interface/controller"
	"gin-sample/backend/interface/gateway"
	"gin-sample/backend/usecase/interactor"
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

		ur := gateway.NewUserRepository(database.GetDB())
		ui := interactor.NewUserInteractor()
		uc := controller.NewUserController(ui)
		userGr.GET("", uc.Get)
		userGr.POST("", uc.Create)
		//userGr.PUT(":id/", uc.Update)
		//userGr.DELETE(":id/", uc.Delete)
	}
}
