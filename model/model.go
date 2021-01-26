package model

import (
	"log"
)

type Userinfo struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
	UserPicture  string `json:"user_picture"`
	//	motto       string `json:"user_motto"`
}

func Register(name string, password string) string {
	//count := 0
	//if err := DB.Table("users").Count(&count).Error; err != nil {
	//	log.Println("registerError" + err.Error())
	//	return ""
	//}
	//temp := 0 + count
	//id := strconv.Itoa(temp)
	user := Userinfo{UserName: name, UserPassword: password}
	if err := DB.Table("users").Create(&user).Error; err != nil {
		log.Println("registError" + err.Error())
		return ""
	}
	DB.Table("users").Create(&user)
	return name
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
	if user.UserPassword == password {
		return true
	}
	log.Println(user.UserPassword)
	log.Println(password)
	return false
}
