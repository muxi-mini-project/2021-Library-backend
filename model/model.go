package model

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

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
		fmt.Println("registError" + err.Error())
		return ""
	}
	//DB.Table("users").Create(&user)
	return name
}

//验证用户是否存在
func IfExistUser(name string) bool {
	var user = make([]Userinfo, 1)
	if err := DB.Table("users").Where("user_name=?", name).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	if len(user) != 0 {
		return true
	}
	return false
}

//验证密码是否正确
func TestPassword(name string, password string) bool {
	var user Userinfo
	if err := DB.Table("users").Where("user_name=? AND user_password=?", name, password).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	if user.UserName != name || user.UserPassword != password {
		return false
	}
	return true
}

//获取userId
func GetId(name string, password string) string {
	var user Userinfo
	if err := DB.Table("users").Where("user_name=? AND user_password=?", name, password).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return " "
	} else {
		return user.UserId
	}
}

type jwtClaims struct {
	jwt.StandardClaims
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
}

var (
	key        = "miniProject" //salt
	ExpireTime = 3600          //token expire time
)

func CreateToken(name string, password string, id string) string {
	//生成token
	//token := jwt.New(jwt.SigningMethodHS256)
	claims := &jwtClaims{
		UserId:   id,
		UserName: name,
		Password: password,
	}
	//token.Claims = claims
	//打印一遍
	//fmt.Println(claims)
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
