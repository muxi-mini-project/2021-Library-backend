package model

import (
	"errors"
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
	return user.UserId
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

var Secret = "sault" //加盐
//验证token
func VerifyToken(token string) (*jwtClaims, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return nil, errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return nil, errors.New("发生错误")
	}
	return claims, nil
}

//获取用户信息
func GetUserName(UserId string) string {
	var user Userinfo
	if err := DB.Table("users").Where("user_id=?", UserId).Find(&user).Error; err != nil {
		return " "
	} else {
		return user.UserName
	}
}

func GetUserPicture(UserId string) string {
	var user Userinfo
	if err := DB.Table("users").Where("user_id=?", UserId).Find(&user).Error; err != nil {
		return " "
	} else {
		return user.UserPicture
	}
}

func GetUserMotto(UserId string) string {
	var user Userinfo
	if err := DB.Table("users").Where("user_id=?", UserId).Find(&user).Error; err != nil {
		return " "
	} else {
		return user.Motto
	}
}

//获取我的书本的id
func GetMyBooksId(UserId string) []func() {
	var bookId []func()
	if err := DB.Table("users_books").Where("user_id=?", UserId).Find(&bookId).Error; err != nil {
		return bookId
	} else {
		return bookId
	}
}

//获取我的书架上书本的详情
func GetMyBooksinfo(BooksId []func()) ([]BooksInfo, error) {
	var Books []BooksInfo
	if err := DB.Table("books").Where("book_id=?", BooksId).Find(&Books).Error; err != nil {
		return nil, err
	}
	return Books, nil
}

//获取我的发布的信息
func GetMyDigest(UserId string) ([]DigestInfo, error) {
	var Digest []DigestInfo
	if err := DB.Table("summaries").Where("user_id=?", UserId).Find(&Digest).Error; err != nil {
		return nil, err
	}
	return Digest, nil
}

//获取用户的信息
func GetUserInfo(UserId string) (Userinfo, error) {
	var user Userinfo
	if err := DB.Table("users").Where("user_id=?", UserId).Find(&user).Error; err != nil {
		return Userinfo{}, err
	}
	return user, nil
}

//修改用户信息
func ChangeUserInfo(user Userinfo) error {
	if err := DB.Table("users").Where("user_id=?", user.UserId).Updates(map[string]interface{}{"user_name": "user.UserName", "user_password": "user.UserPassword", "user_picture": "user.UserPicture", "motto": "Motto"}).Error; err != nil {
		return err
	}
	return nil
}
