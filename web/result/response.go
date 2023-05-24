package result

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Success(c *gin.Context, code int, data any) {

	c.PureJSON(200, Response{
		Code:    code,
		Data:    data,
		Message: "success",
	})
}

func Fail(c *gin.Context, code int, err error) {

	c.PureJSON(200, Response{
		Code:    code,
		Data:    nil,
		Message: err.Error(),
	})
}
