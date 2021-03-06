package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func Initdb() *gorm.DB {
	DB, err = gorm.Open("mysql", "jingzhe:123@tcp(0.0.0.0:3306)/study?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
	DB.LogMode(true)
	return DB
}
