package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/KouT127/gin-sample/backend/config"
	"github.com/KouT127/gin-sample/backend/interface/graphql/generated"
	"github.com/KouT127/gin-sample/backend/interface/middlewares"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{"authorization", "Content-Type"},
	}))
	e.GET("/", playgroundHandler())
	graph := e.Group("/query", middlewares.FirebaseAuth())
	{
		graph.Use(dataloader.LoaderMiddleware())
		graph.POST("", graphqlHandler(c))
	}
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
		handler.ComplexityLimit(200),
	))
}
func playgroundHandler() echo.HandlerFunc {
	return echo.WrapHandler(handler.Playground("GraphQL", "/query"))
}
