package http

import (
	"net/http"

	docs "github.com/joqd/slovo/cmd/docs"
	"github.com/joqd/slovo/internal/adapter/config"
	v1 "github.com/joqd/slovo/internal/adapter/delivery/http/v1"
	"github.com/joqd/slovo/internal/core/port"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RouterOptions struct {
	Engine *gin.Engine
	Conf   *config.Config
	Log    port.Logger

	WordUsecase port.WordUsecase
}

func RegisterRoutes(opts RouterOptions) {
	// Swagger
	if opts.Conf.Swagger.Enabled {
		docs.SwaggerInfo.BasePath = "/"
		opts.Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	// Load html and static files
	opts.Engine.LoadHTMLGlob("web/templates/*.html")
	opts.Engine.Static("/static", "./web/static")

	// Ping check
	opts.Engine.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ping.html", gin.H{
			"version": opts.Conf.App.Version,
		})
	})

	// Routers
	apiV1Group := opts.Engine.Group("/api/v1")
	{
		v1.RegisterWordRouter(apiV1Group, opts.WordUsecase, opts.Log)
	}
}
