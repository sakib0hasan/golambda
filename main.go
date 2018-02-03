package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/walcott911/golambda/config"
)

func main() {
	fmt.Println("GoLambda Running")

	router := gin.Default()

	// Web frontend
	console := router.Group("/console")
	{
		console.GET("/", InfoHandler)
		console.GET("/login", InfoHandler)
	}

	api := router.Group("/api")
	{
		api.GET("/info", InfoHandler)
		api.GET("/login", InfoHandler)
		api.GET("/tenant/register", InfoHandler)
		api.GET("/lambda/get/:id", InfoHandler)
		api.GET("/lambda/create/:id", InfoHandler)
		api.GET("/lambda/update/:id", InfoHandler)
		api.GET("/lambda/delete/:id", InfoHandler)
	}
	router.Run(":8080")
}

func InfoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    config.APP_NAME,
		"Version": config.VERSION,
	})
}
