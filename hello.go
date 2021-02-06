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
		header := c.Request.Header["HTTP_X_GITHUB_EVENT"]
		fmt.Println(header)
		payload := c.PostForm("payload")

		f, err := os.Create("data.json")

		if err != nil {
			fmt.Println("error")
		} else {
			f.WriteString(payload)
		}

		fmt.Println("saved")
	})

	router.Run(":8080")
}
