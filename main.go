package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	routes "github.com/gopher-fleece/gleecexample/routers"
)

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

	// Start the server on port 8080
	router.Run("127.0.0.1:8080")
}
