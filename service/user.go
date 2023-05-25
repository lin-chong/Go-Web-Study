package service

import (
	"Go-Web-Study/db"
	"Go-Web-Study/service/payload"
	"Go-Web-Study/web/result"
	"github.com/gin-gonic/gin"
	"log"
)

func GetUserById(ctx *gin.Context) {
	var req payload.GetUserByIdReq
	if err := ctx.ShouldBindQuery(&req); err != nil {
		result.Fail(ctx, 500, err)
		return
	}
	log.Printf("id: %d", req.Id)

	conn := db.GetDBInstance()

	user, err := conn.GetUserByUserId(ctx, req.Id)

	if err != nil {
		result.Fail(ctx, 500, err)
		return
	}

	result.Success(ctx, 200, user)
}
