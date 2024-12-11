package main

import (

	"fmt"
	"github.com/gin-gonic/gin"
	"server/config"

)

func main() {
	// Load the MongoDB secret
	mongodbURI := config.MongoDBSecret()

	// Connect to the MongoDB database
	config.ConnectDB(mongodbURI)

	// Set the router as the default one provided by Gin
	router := gin.Default()

	// Serve the front end 
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run the server
	router.Run(":8080")
	fmt.Println("Server is running on port 8080")

}