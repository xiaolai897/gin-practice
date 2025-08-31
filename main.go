package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)
	router.Run("0.0.0.0:7200")
}
