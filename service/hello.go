package service

import (
	"Go-Web-Study/web/result"
	"github.com/gin-gonic/gin"
	"log"
)

func Hello(c *gin.Context) {
	log.Println("Hello World")
	result.Success(c, 200, "Hello World")
}
