package server

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/KouT127/gin-sample/backend/infrastracture/database"
	"github.com/KouT127/gin-sample/backend/infrastracture/graphql"
	"github.com/KouT127/gin-sample/backend/infrastracture/handlers"
	"github.com/KouT127/gin-sample/backend/infrastracture/middlewares"
	"github.com/KouT127/gin-sample/backend/interface/controller"
	"github.com/KouT127/gin-sample/backend/interface/gateway"
	"github.com/KouT127/gin-sample/backend/interface/presenter"
	"github.com/KouT127/gin-sample/backend/usecase/interactor"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/uber/jaeger-lib/metrics"

	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
)

var db *gorm.DB

func NewRouter() *gin.Engine {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	// Initialize tracer with a logger and a metrics factory
	closer, _ := cfg.InitGlobalTracer(
		"greetings-server",
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	defer closer.Close()
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
