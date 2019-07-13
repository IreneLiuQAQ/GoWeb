package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/echo", postEcho)
	err := router.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}

func postEcho(c *gin.Context) {
	text1 := c.DefaultPostForm("content1", "")
	text2 := c.DefaultPostForm("content2", "")
	c.JSON(200, gin.H{
		"string": text1 + text2,
		"name": "meimei",
	})
}
