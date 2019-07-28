package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/KouT127/gin-sample/backend/config"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/dataloader"
	"github.com/KouT127/gin-sample/backend/interface/resolver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := NewRouter()
	e.Logger.Fatal(e.Start(":8080"))
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(config.DebugLoggerConfig))
	e.Use(middleware.CORS())
	e.Use(dataloader.LoaderMiddleware())
	e.POST("/query", graphqlHandler())
	e.GET("/", playgroundHandler())
	return e
}

func graphqlHandler() echo.HandlerFunc {
	return echo.WrapHandler(handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))
}
func playgroundHandler() echo.HandlerFunc {
	return echo.WrapHandler(handler.Playground("GraphQL", "/query"))
}
