package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run(":7001")
	if err != nil {
		return
	} // listen and serve on
}
