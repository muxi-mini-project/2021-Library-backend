package model

import (
	"log"
	"strconv"

	"Library/model"
)

type Userinfo struct {
	Id          string
	UserName    string
	Password    string
	Userpicture string
	motto       string
}

func Register(name string, password string) string {
	count := 0
	if err := model.DB.Table("user").Count(&count).Error; err != nil {
		log.Println("registerError" + err.Error())
		return ""
	}
	temp := 0 + count
	id := strconv.Itoa(temp)
	user := Userinfo{Id: id, UserName: name, Password: password}
	if err := model.DB.Table("user").Create(&user).Error; err != nil {
		log.Println("registError" + err.Error())
		return ""
	}
	return ID
}
