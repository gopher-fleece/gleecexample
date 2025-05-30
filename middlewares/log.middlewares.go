package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
)

func LogBeforeOperationMiddleware(ctx context.Context, ginCtx *gin.Context) (context.Context, bool) {
	println("Method: ", ginCtx.Request.Method, " Path: ", ginCtx.Request.URL.Path, " arrived")
	// An example of loading contextual data on the context
	ctx = context.WithValue(ctx, "user", ginCtx.GetHeader("user"))
	return ctx, true
}

func LogAfterOperationSuccessMiddleware(ctx context.Context, ginCtx *gin.Context) (context.Context, bool) {
	println("Method: ", ginCtx.Request.Method, " Path: ", ginCtx.Request.URL.Path, " completed")
	return ctx, true
}

func LogOnErrorMiddleware(ctx context.Context, ginCtx *gin.Context, err error) (context.Context, bool) {
	println("Method: ", ginCtx.Request.Method, " Path: ", ginCtx.Request.URL.Path, " failed with error: ", err.Error())
	return ctx, true
}

func LogOnValidationErrorMiddleware(ctx context.Context, ginCtx *gin.Context, err error) (context.Context, bool) {
	println("Method: ", ginCtx.Request.Method, " Path: ", ginCtx.Request.URL.Path, " input validation failed: ", err.Error())
	return ctx, true
}
