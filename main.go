package main

import (
	"demo-project/config"
	"demo-project/models"
	"demo-project/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config.Connect()

	// Run migrations
	models.MigrateUser(config.GetDB())
	models.MigrateStudent()

	// Initialize the Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})
	// Start the server
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
