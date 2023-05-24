package service

import (
	"Go-Web-Study/db"
	"Go-Web-Study/web/result"
	"github.com/gin-gonic/gin"
	"log"
)

func Hello(c *gin.Context) {
	log.Println("Hello World")
	result.Success(c, 200, "Hello World")
}

func Dataset(c *gin.Context) {
	conn := db.GetDB()

	query, err := conn.Query("select `title` from `datasets` where id = ?", 1)

	if err != nil {
		result.Fail(c, 500, err)
	}

	var title string
	if query.Next() {

		err := query.Scan(&title)

		if err != nil {
			result.Fail(c, 500, err)
		}
	}

	result.Success(c, 200, title)

}
