package respose

import "github.com/gin-gonic/gin"

func Sucess(data interface{}) gin.H{
	return gin.H{
		"code":200,
		"msg" :"sucess",
		"data" :data,
	}
}
func Error() gin.H{
	return gin.H{
		"code":500,
		"msg" :"error",
	}
}
func ErrorWithMsg(msg string) gin.H{
	return gin.H{
		"code":500,
		"msg" :msg,
	}
}
func NewResult(code int, msg string,data interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}