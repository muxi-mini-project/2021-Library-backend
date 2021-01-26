package main

import (
	"2021-Library-backend/model"
	"log"

	"2021-Library-backend/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()

	DB := model.Initdb()
	defer DB.Close()

	router.Router(r)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}

}
