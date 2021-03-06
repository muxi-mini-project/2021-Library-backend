package main

import (
	"2021-Library-backend/model"
	"log"

	"2021-Library-backend/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// @title 2021-miniproject-Library
// @version 1.0
// @description This is a project made by CCNU 书斋

// @BaesPath /Library/v1/

// @Schemas http

func main() {
	r := gin.Default()

	model.DB = model.Initdb()
	defer model.DB.Close()

	router.Router(r)
	if err := r.Run(":10086"); err != nil {
		log.Fatal(err.Error())
	}

}
