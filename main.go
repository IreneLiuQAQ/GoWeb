package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

type User struct {
	Username string `gorm:"primary_key"`
	Nickname string
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./user.db")
	if err != nil {
		panic(err)
	}
	db = db.Debug()
	db.AutoMigrate(User{})
}

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"string": "Hello World!",
		})
	})
	router.POST("/echo", postEcho)
	router.POST("/calc/sum", postCalc)
	router.POST("/user/register", postUser)
	router.GET("/user/:username", getNickname)
	err := router.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}

func postUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	nickname := c.DefaultPostForm("nickname", "")
	user1 := User{Username: username, Nickname: nickname}
	db.Create(&user1)
	var err error
	var user User
	var code int
	if err = db.Where("`username` = ?", username).Find(&user).Error; err != nil {
		code = http.StatusForbidden
	} else {
		code = http.StatusCreated
	}
	c.JSON(200, gin.H{
		"code":    code,
		"massage": "Created success",
	})
}

func getNickname(c *gin.Context) {
	username := c.Param("username")
	var user User
	db.Where("username = ?", username).Find(&user)
	nickname := user.Nickname
	c.JSON(200, gin.H{
		"nickname": nickname,
		"code":     http.StatusOK,
		"massage":  "success",
	})
}

func postEcho(c *gin.Context) {
	text1 := c.DefaultPostForm("content1", "")
	text2 := c.DefaultPostForm("content2", "")
	c.JSON(200, gin.H{
		"string": text1 + text2,
		"name": "" +
			"meimei",
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
