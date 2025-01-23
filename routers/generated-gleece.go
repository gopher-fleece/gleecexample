package routes
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gopher-fleece/gleece/external"
	UsersControllerImport "github.com/gopher-fleece/gleecexample/controllers"
)
var validatorInstance *validator.Validate
var urlParamRegex = regexp.MustCompile(`\{([\w\d-_]+)\}`)
func getStatusCode(controller external.Controller, hasReturnValue bool, err error) int {
	if controller.GetStatus() != nil {
		return int(*controller.GetStatus())
	}
	if err != nil {
		return http.StatusInternalServerError
	}
	if hasReturnValue {
		return http.StatusOK
	}
	return http.StatusNoContent
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
func toGinUrl(url string) string {
	return urlParamRegex.ReplaceAllString(url, ":$1")
}
func RegisterRoutes(engine *gin.Engine) {
	validatorInstance = validator.New()
	// UsersController
	engine.POST(toGinUrl("/users/user/{user_name}"), func(ctx *gin.Context) {
		controller := UsersControllerImport.UsersController{}
		controller.SetRequest(ctx)
		emailRaw := ctx.Query("email")
		email := emailRaw
		nameRaw := ctx.Param("name")
		name := nameRaw
		originRaw := ctx.GetHeader("origin")
		origin := originRaw
		traceRaw := ctx.GetHeader("trace")
		trace := traceRaw
		value, opError := controller.CreateNewUser(email, name, origin, trace)
		statusCode := getStatusCode(&controller, true, opError)
		ctx.Header("Content-Type", "application/json")
		if opError != nil {
			stdError := external.Rfc7807Error{
				Type:       http.StatusText(statusCode),
				Detail:     "Encountered an error during operation 'CreateNewUser'",
				Status:     statusCode,
				Instance:   "/gleece//CreateNewUser",
				Extensions: map[string]string{"error": opError.Error()},
			}
			ctx.JSON(statusCode, stdError)
			return
		}
		ctx.JSON(statusCode, value)
	})
	engine.GET(toGinUrl("/users/domicile/{id}"), func(ctx *gin.Context) {
		controller := UsersControllerImport.UsersController{}
		controller.SetRequest(ctx)
		idRaw := ctx.Param("id")
		id := idRaw
		value, opError := controller.GetUserDomicile(id)
		statusCode := getStatusCode(&controller, true, opError)
		ctx.Header("Content-Type", "application/json")
		if opError != nil {
			stdError := external.Rfc7807Error{
				Type:       http.StatusText(statusCode),
				Detail:     "Encountered an error during operation 'GetUserDomicile'",
				Status:     statusCode,
				Instance:   "/gleece//GetUserDomicile",
				Extensions: map[string]string{"error": opError.Error()},
			}
			ctx.JSON(statusCode, stdError)
			return
		}
		ctx.JSON(statusCode, value)
	})
}
