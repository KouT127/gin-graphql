package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/KouT127/gin-sample/backend/interface/graphql"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/dataloader"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/security"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var db *gorm.DB

//func Init() {
//	r := NewRouter()
//	r.Run()
//}

func Init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(dataloader.LoaderMiddleware)
	e.POST("/query", echo.WrapHandler(handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))))
	e.GET("/", echo.WrapHandler(handler.Playground("GraphQL", "/query")))
	e.Start(":8080")
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(security.CORSMiddleware())
	//r.Use(dataloader.LoaderMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	v1 := r.Group("v1")
	{
		newUserHandler(v1)
	}
	return r
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
