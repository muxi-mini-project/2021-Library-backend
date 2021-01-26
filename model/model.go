package model

import (
	"log"
	"strconv"
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
	if err := DB.Table("user").Count(&count).Error; err != nil {
		log.Println("registerError" + err.Error())
		return ""
	}
	temp := 0 + count
	id := strconv.Itoa(temp)
	user := Userinfo{Id: id, UserName: name, Password: password}
	if err := DB.Table("user").Create(&user).Error; err != nil {
		log.Println("registError" + err.Error())
		return ""
	}
	return id
}

//验证用户是否存在
func IfExistUser(id string) bool {
	var user = make([]Userinfo, 1)
	if err := DB.Table("users").Where("user_id=?", id).First(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	if len(user) != 0 {
		return true
	}
	return false
}

//验证密码是否正确
func TestPassword(id string, password string) bool {
	var user Userinfo
	if err := DB.Table("users").Where("user_id=?", id).Error; err != nil {
		log.Println("Error" + err.Error())
		return false
	}
	if user.Password == password {
		return true
	}
	log.Println(user.Password)
	log.Println(password)
	return false
}
