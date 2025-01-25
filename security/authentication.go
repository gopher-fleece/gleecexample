package security

import (
	"github.com/gin-gonic/gin"
	"github.com/gopher-fleece/gleece/external"
)

func GleeceRequestAuthorization(ctx *gin.Context, check external.SecurityCheck) *external.SecurityError {

	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "change that condition...." {
		return &external.SecurityError{
			StatusCode: 403,
			Message:    "You are not authorized to read that API",
		}
	}
	return nil
}
