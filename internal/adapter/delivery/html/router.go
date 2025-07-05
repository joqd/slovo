package html

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/joqd/slovo/internal/adapter/config"
	"github.com/joqd/slovo/internal/adapter/delivery/html/handler"
	"github.com/joqd/slovo/internal/core/port"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Engine *gin.Engine
	Conf   *config.Config
	Log    port.Logger
}

func RegisterRoutes(opts Options) {
	opts.Engine.HTMLRender = loadTemplates("./internal/adapter/delivery/html/templates")
	opts.Engine.Static("/static", "./internal/adapter/delivery/html/static")

	router := opts.Engine.Group("/")

	{
		handler.RegisterIndexRouter(router, opts.Log)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		layoutCopy = append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), layoutCopy...)
	}

	return r
}
