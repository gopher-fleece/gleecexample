package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
	routes "github.com/gopher-fleece/gleecexample/routes"
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
