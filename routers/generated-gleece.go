package routes
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gopher-fleece/gleece/external"
)
var validatorInstance *validator.Validate
func getStatusCode(controller external.Controller, hasReturnValue bool, err error) int {
	if controller.GetStatus() != nil {
		return int(*controller.GetStatus())
	}
	if err != nil {
		return http.StatusInternalServerError
	}
	if hasReturnValue {
		return http.StatusNoContent
	}
	return http.StatusOK
}
func bindAndValidateBody[TOutput any](ctx *gin.Context, contentType string, output *TOutput) error {
	var err error
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return err
	}
	var deserializedOutput TOutput
	switch contentType {
	case "application/json":
		err = json.Unmarshal(bodyBytes, &deserializedOutput)
	default:
		return fmt.Errorf("content-type %s is not currently supported by the validation subsystem", contentType)
	}
	if err != nil {
		return err
	}
	if err = validatorInstance.Struct(&deserializedOutput); err != nil {
		return err
	}
	output = &deserializedOutput
	return nil
}
func RegisterRoutes(engine *gin.Engine) {
	validatorInstance = validator.New()
}
