package web

import (
	"Go-Web-Study/service"
	"github.com/gin-gonic/gin"
)

func Route(c *gin.Engine) {
	c.Group("/api").
		GET("/hello", service.Hello)
}
