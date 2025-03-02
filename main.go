package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gopher-fleece/gleecexample/middlewares"
	routes "github.com/gopher-fleece/gleecexample/routes"
	"github.com/gopher-fleece/gleecexample/validators"
	"github.com/gopher-fleece/runtime"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed openapi/swagger.json
var swaggerSpec embed.FS

// Handler to serve the Swagger spec file
func serveSwaggerSpec(c *gin.Context) {
	// Read the embedded swagger.json file
	specData, err := swaggerSpec.ReadFile("openapi/swagger.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read swagger spec",
		})
		return
	}

	// Set the content type to application/json
	c.Header("Content-Type", "application/json")
	// Write the spec data directly to the response
	c.Writer.Write(specData)
}

func main() {

	// Create a default gin router
	router := gin.Default()

	// Define a route for GET /hello using the native engine
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// # Gleece integration part

	// Register custom validation rules
	routes.RegisterCustomValidator("validate_starts_with_letter", validators.ValidateStartsWithLetter)
	// Register middlewares
	routes.RegisterMiddleware(runtime.BeforeOperation, middlewares.LogBeforeOperationMiddleware)
	routes.RegisterMiddleware(runtime.AfterOperationSuccess, middlewares.LogAfterOperationSuccessMiddleware)
	routes.RegisterErrorMiddleware(runtime.OnOperationError, middlewares.LogOnErrorMiddleware)
	routes.RegisterErrorMiddleware(runtime.OnInputValidationError, middlewares.LogOnValidationErrorMiddleware)
	// Register the the Gleece generated routes
	routes.RegisterRoutes(router)

	// # End Gleece integration part

	// Serve the Swagger spec file at /openapi/swagger.json
	router.GET("/openapi/swagger.json", serveSwaggerSpec)

	// Serve the Swagger UI at /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/openapi/swagger.json")))

	// Start the server on port 8080
	router.Run("127.0.0.1:8080")
}
