package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/event_handler", func(c *gin.Context) {
		payload := c.PostForm("payload")

		f, err := os.Create("data.txt")

		if err != nil {
			f.WriteString(payload)
		}

		fmt.Printf(payload)
	})

	router.Run(":8080")
}
