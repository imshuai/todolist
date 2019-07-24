package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tpl", config)
	return
}
