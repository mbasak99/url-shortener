package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener!",
		})
	})

	err := r.Run(":3000")
	if err != nil {
		fmt.Errorf("Failed to start the web server - Error: %v\n", err)
	}
}
