package security

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gopher-fleece/runtime"
)

func GleeceRequestAuthorization(ctx context.Context, ginCtx *gin.Context, check runtime.SecurityCheck) (context.Context, *runtime.SecurityError) {
	authHeader := ginCtx.GetHeader("Authorization")
	if authHeader == "change that condition...." {
		return ctx, &runtime.SecurityError{
			StatusCode: 403,
			Message:    "You are not authorized to read that API",
		}
	}
	return ctx, nil
}
