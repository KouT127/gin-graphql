package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/KouT127/gin-sample/backend/config"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	Authenticaiton "github.com/KouT127/gin-sample/backend/interface/middlewares/authentication"
	"github.com/KouT127/gin-sample/backend/interface/middlewares/dataloader"
	"github.com/KouT127/gin-sample/backend/interface/resolver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() {
	e := NewRouter()
	e.Logger.Fatal(e.Start(":8080"))
}

func NewRouter() *echo.Echo {

	e := echo.New()
	c := complexityHandler()
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(config.DebugLoggerConfig))
	e.Use(Authenticaiton.FirebaseAuth())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{""},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{"Content-Type"},
	}))

	e.Use(dataloader.LoaderMiddleware())
	e.POST("/query", graphqlHandler(c))
	e.GET("/", playgroundHandler())
	return e
}

func complexityHandler() *generated.Config {
	c := generated.Config{Resolvers: &resolver.Resolver{}}
	connectionComplexity := func(childComplexity int, first *int, after *string, last *int, before *string, query *string) int {
		if first != nil {
			return *first * childComplexity
		} else if last != nil {
			return *last * childComplexity
		}
		return childComplexity
	}
	c.Complexity.Query.Tasks = connectionComplexity
	return &c
}

func graphqlHandler(c *generated.Config) echo.HandlerFunc {
	return echo.WrapHandler(handler.GraphQL(
		generated.NewExecutableSchema(*c),
		handler.ComplexityLimit(5),
	))
}
func playgroundHandler() echo.HandlerFunc {
	return echo.WrapHandler(handler.Playground("GraphQL", "/query"))
}
