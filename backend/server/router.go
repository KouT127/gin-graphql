package server

import (
	"github.com/99designs/gqlgen/handler"
	. "github.com/KouT127/gin-sample/backend/config"
	"github.com/KouT127/gin-sample/backend/interface/graphql"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/dataloader"
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
	e := NewRouter()
	e.Start(":8080")
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(DebugLoggerConfig))
	e.Use(dataloader.LoaderMiddleware)
	e.POST("/query", graphqlHandler())
	e.GET("/", playgroundHandler())
	return e
}

func graphqlHandler() echo.HandlerFunc {
	return echo.WrapHandler(handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}})))
}
func playgroundHandler() echo.HandlerFunc {
	return echo.WrapHandler(handler.Playground("GraphQL", "/query"))
}
