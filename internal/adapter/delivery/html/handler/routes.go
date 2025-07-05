package handler

import (
	"github.com/joqd/slovo/internal/core/port"

	"github.com/gin-gonic/gin"
)

func RegisterIndexRouter(rg *gin.RouterGroup, xlog port.Logger) {
	indexHandler := NewIndexHandler(xlog)

	{
		rg.GET("/", indexHandler.Index)
	}
}
