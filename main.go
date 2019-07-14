package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"strconv"
)

type Person struct {
	ID        uint
	FirstName string
	LastName  string
}

type User struct {
	Username string
	Nickname string
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./gorm.db")
	db, err = gorm.Open("sqlite3","./user.db")
	if err != nil {
		panic(err)
	}
}

func gromPerson()  {
	db.AutoMigrate(&Person{})
	p1 := Person{FirstName:"Qinagyuan", LastName:"Shui"}
	db.Create(&p1)
}

func gromUser(){
	db.AutoMigrate(&User{})
	u1 := User{Username:"username",Nickname:"nickname"}
	db.Create(&u1)

}
func main() {
	gromPerson()
	gromUser()
	router := gin.Default()
	router.POST("/echo", postEcho)
	router.POST("/calc/sum", postCalc)
	router.GET("/db", GetProjects)
	router.POST("/user/register",Postuser)
	err := router.Run("0.0.0.0:8000")
	if err != nil {
		panic(err)
	}
}

func GetProjects(con *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		con.AbortWithStatus(404)
		fmt.Println(err)
	} else{
		con.JSON(200,people)
	}

}
func Postuser(c *gin.Context){
	c.JSON(200, gin.H{
		"code": http.StatusCreated,
		"massage": "Created success",
	})
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
