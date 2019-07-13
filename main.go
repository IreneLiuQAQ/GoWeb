package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	router := gin.Default()
	router.POST("/echo", postEcho)
	router.POST("/calc/sum", postCalc)
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
		"name":   "meimei",
	})
}

func postCalc(context *gin.Context) {
	num1 := context.DefaultPostForm("a", "0")
	num2 := context.DefaultPostForm("b", "0")
	context.JSON(200, gin.H{
		"result": calc(num1, num2),
	})
}

func calc(a string, b string) (int64) {
	num1, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		panic(err)
	}
	return num1 + num2
}
