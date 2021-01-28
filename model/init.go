package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB(DB *gorm.DB) *gorm.DB{
	dns := "root:********@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Summary{})
	DB.AutoMigrate(&SummaryClass{})
	DB.AutoMigrate(&Book{})
	return DB
}


