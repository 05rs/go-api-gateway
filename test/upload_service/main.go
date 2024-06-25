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
	router.GET("/upload/health", health)
	router.GET("/upload/files/{id} ", getFile)

	// Start the server
	router.Run(":8081")
}

func health(c *gin.Context) {
	fmt.Println("PING SUCCESS!!!")
	c.String(http.StatusOK, "OK")
}

func getFile(c *gin.Context) {
	fmt.Println("getFile Request: id:", c.Param("id"))
	c.String(http.StatusOK, "File id:", c.Param("id"))
}
