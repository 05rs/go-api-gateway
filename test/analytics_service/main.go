package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin router
	router := gin.Default()

	// Register Routes
	router.GET("/analytics/health", health)

	// Start the server
	router.Run(":8082")
}

func health(c *gin.Context) {
	fmt.Println("PING SUCCESS!!!")
	c.String(http.StatusOK, "OK")
}
