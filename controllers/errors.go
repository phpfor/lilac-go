package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceError(c *gin.Context) {
	c.HTML(http.StatusNotFound, "errors/500", nil)
}

//NotFound handles gin NotFound error
func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "errors/404", nil)
}

//MethodNotAllowed handles gin MethodNotAllowed error
func MethodNotAllowed(c *gin.Context) {
	c.HTML(http.StatusMethodNotAllowed, "errors/405", nil)
}
