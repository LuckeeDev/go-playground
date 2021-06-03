package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()

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
		
		// SECRET_TOKEN extracted from the .env file
		var token string = os.Getenv("SECRET_TOKEN")	

		// The body content
		var data []byte
		if c.Request.Body != nil {
			data, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(200, err)
		}

		if err := c.ShouldBindHeader(&headers); err != nil {
			c.JSON(200, err)
		}
		
		result, validationError := validateSignature(token, data, headers.GithubSignature)

		if validationError != nil {
			c.JSON(200, err)
		}

		if result == true {
			c.JSON(200, gin.H{ "header": "Valid request." })
		} else {
			c.JSON(403, gin.H{ "result": "Seems like you're not authorized to perform this action!" })
		}
	})

	router.Run(":8080")
}
