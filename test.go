package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	//"net/http"
)

type users_books struct {
	Id int64
	User_id int64 `json:"ui"`
	Book_id int64 `json:"bi"`
}

func main() {
	db, err := gorm.Open("mysql", "root:hu20010326@(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()

	r := gin.Default()
	r.POST("/", func(c *gin.Context){
		var auser users_books
		err = c.BindJSON(&auser)
		if err != nil {
			c.JSON(200, gin.H{
				"errcode": 400,
				"description": "Post Data Err",
			})
		return
		} else {
			fmt.Println(auser.User_id)
			db.AutoMigrate(&users_books{})
			db.Create(&auser)
		}
	})
	r.GET("/", func(c *gin.Context) {
		var auser []users_books
		db.Find(&auser)
		c.JSON(200, gin.H{
			"data": auser,
		})
		//fmt.Println(auser.User_id)
		//name := c.Query("name")
		//role := c.DefaultQuery("role", "teacher")
		//c.String(200, "%s is a %s", name, role)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
