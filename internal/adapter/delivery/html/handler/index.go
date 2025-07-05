package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joqd/slovo/internal/core/port"
)

type indexHandler struct {
	xlog port.Logger
}

func NewIndexHandler(xlog port.Logger) *indexHandler {
	return &indexHandler{
		xlog: xlog,
	}
}

func (i *indexHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Index",
	})
}
