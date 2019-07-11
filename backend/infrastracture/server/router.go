package server

import (
	"gin-sample/backend/infrastracture/database"
	"gin-sample/backend/infrastracture/graphql"
	"gin-sample/backend/infrastracture/handlers"
	"gin-sample/backend/infrastracture/middlewares"
	"gin-sample/backend/interface/controller"
	"gin-sample/backend/interface/gateway"
	"gin-sample/backend/interface/presenter"
	"gin-sample/backend/usecase/interactor"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.CORSMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	newHealthHandler(r)
	v1 := r.Group("v1")
	{
		newUserHandler(v1)
	}
	return r
}

func newHealthHandler(router *gin.Engine) {
	health := new(handlers.HealthHandler)
	router.GET("health", health.Status)
}

func newUserHandler(gr *gin.RouterGroup) {
	userGr := gr.Group("users")
	{
		ur := gateway.NewUserRepository(database.GetDB())
		up := presenter.NewUserPresenter()
		ui := interactor.NewUserInteractor(ur, up)
		uc := controller.NewUserController(ui)
		userGr.GET("", uc.Get)
		userGr.POST("", uc.Create)
		//userGr.PUT(":id/", uc.Update)
		//userGr.DELETE(":id/", uc.Delete)
	}
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}