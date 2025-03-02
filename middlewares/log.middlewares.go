package middlewares

import (
	"github.com/gin-gonic/gin"
)

func LogBeforeOperationMiddleware(ctx *gin.Context) bool {
	println("Method: ", ctx.Request.Method, " Path: ", ctx.Request.URL.Path, " arrived")
	return true
}

func LogAfterOperationSuccessMiddleware(ctx *gin.Context) bool {
	println("Method: ", ctx.Request.Method, " Path: ", ctx.Request.URL.Path, " completed")
	return true
}

func LogOnErrorMiddleware(ctx *gin.Context, err error) bool {
	println("Method: ", ctx.Request.Method, " Path: ", ctx.Request.URL.Path, " failed with error: ", err.Error())
	return true
}

func LogOnValidationErrorMiddleware(ctx *gin.Context, err error) bool {
	println("Method: ", ctx.Request.Method, " Path: ", ctx.Request.URL.Path, " input validation failed: ", err.Error())
	return true
}
