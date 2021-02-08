package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()

	SecretToken := os.Getenv("SECRET_TOKEN")

	fmt.Println(SecretToken)

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/event_handler", func(c *gin.Context) {
		var headers *Headers
		var form *Form

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(200, err)
		}

		if err := c.ShouldBindHeader(&headers); err != nil {
			c.JSON(200, err)
		}
	
		fmt.Printf("%#v\n", headers)
		c.JSON(200, gin.H{ "payload": form.Payload, "header": headers.GithubSignature })
	})

	router.Run(":8080")
}
