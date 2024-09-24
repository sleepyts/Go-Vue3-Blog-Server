package respose

import "github.com/gin-gonic/gin"

func Sucess(data interface{}) gin.H {
	return gin.H{
		"code":    200,
		"message": "sucess",
		"data":    data,
	}
}
func Error() gin.H {
	return gin.H{
		"code":    500,
		"message": "error",
	}
}
func ErrorWithMsg(msg string) gin.H {
	return gin.H{
		"code":    500,
		"message": msg,
	}
}
func NewResult(code int, msg string, data interface{}) gin.H {
	return gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}
