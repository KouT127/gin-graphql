package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/KouT127/gin-sample/backend/interface/graphql"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/handlers"
	"github.com/KouT127/gin-sample/backend/interface/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func NewRouter() *gin.Engine {
	r := gin.New()
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
		uc := InjectUser()
		userGr.GET("", uc.Get)
		userGr.POST("", uc.Create)
		//userGr.PUT(":id/", uc.Update)
		//userGr.DELETE(":id/", uc.Delete)
	}
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

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
