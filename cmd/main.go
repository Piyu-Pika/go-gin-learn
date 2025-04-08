package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.New()
	logrus.Println("Starting server on port 8080")

	router.GET("/", hello)

	router.Run(":8080")

}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
