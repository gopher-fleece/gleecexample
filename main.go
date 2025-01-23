package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gleece "github.com/gopher-fleece/gleece/cmd"
	"github.com/gopher-fleece/gleece/cmd/arguments"
	"github.com/gopher-fleece/gleece/infrastructure/validation"
	routes "github.com/gopher-fleece/gleecexample/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Handler to serve the Swagger spec file
func serveSwaggerSpec(c *gin.Context) {
	c.File("openapi/swagger.json")
}

func buildApiProgrammatically() {
	// TODO: expose official Gleece API to build the API programmatically
	validation.InitValidator()
	gleece.GenerateSpecAndRoutes(arguments.CliArguments{ConfigPath: "./gleece.json"})
}

func main() {
	if false {
		buildApiProgrammatically()
	} else {

		// Create a default gin router
		router := gin.Default()

		// Define a route for GET /hello
		router.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World!",
			})
		})

		// Register the the Gleece generated routes
		routes.RegisterRoutes(router)

		// Serve the Swagger spec file at /openapi/swagger.json
		router.GET("/openapi/swagger.json", serveSwaggerSpec)

		// Serve the Swagger UI at /swagger/index.html
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/openapi/swagger.json")))

		// Start the server on port 8080
		router.Run("127.0.0.1:8080")
	}
}
