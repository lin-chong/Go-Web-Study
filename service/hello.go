package service

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Hello(c *gin.Context) {
	log.Println("Hello World")
	c.PureJSON(200, "Hello World")
}
