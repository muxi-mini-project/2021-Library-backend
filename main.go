package main

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := gin.Default()

	DB := model.Initdb()
	defer DB.Close()

	router.Run()

}
