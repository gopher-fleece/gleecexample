package security

import (
	"github.com/gin-gonic/gin"
	"github.com/gopher-fleece/runtime"
)

func GleeceRequestAuthorization(ctx *gin.Context, check runtime.SecurityCheck) *runtime.SecurityError {

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "change that condition...." {
		return &runtime.SecurityError{
			StatusCode: 403,
			Message:    "You are not authorized to read that API",
		}
	}
	return nil
}
