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
		return " "
	}
	//DB.Table("users").Create(&user)
	return user.UserId
}

//验证用户是否存在
func IfExistUser(name string) bool {
	var user = make([]Userinfo, 1)
	if err := DB.Table("users").Where("user_name=?", name).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	if len(user) != 1 {
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

//生成token与验证
type jwtClaims struct {
	jwt.StandardClaims
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
}

/*func GetToken(claims *JwtClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		log.Println(err)
		return ""
	}
	return signedToken
}
*/

var (
	key        = "miniProject" //salt
	ExpireTime = 604800        //token expire time
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

//var key = "sault" //加盐
//验证token
func VerifyToken(token string) (string, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", errors.New("token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return "", errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return "", errors.New("发生错误")
	}
	return claims.UserId, nil
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

type temp struct {
	Id     int64  `gorm:"column:id"`
	UserId string `gorm:"column:user_id"`
	BookId string `gorm:"column:book_id"`
}

//获取我的书本的id
func GetMyBooksId(UserId string) []string {
	var Id []string
	var BookId []temp
	//fmt.Println(UserId)
	if err := DB.Table("users_books").Where("user_id=?", UserId).Find(&BookId).Error; err != nil {
		log.Println(err)
		return nil
	} else {
		fmt.Println(BookId)
		for _, id := range BookId {
			Id = append(Id, string(id.BookId))
			//fmt.Println(id.BookId)
		}
		return Id
	}
}

//获取我的书架上书本的详情
func GetMyBooksinfo(BooksId []string) ([]BooksInfo, error) {
	var Books []BooksInfo
	if err := DB.Table("books").Where("book_id in (?)", BooksId).Find(&Books).Error; err != nil {
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
	fmt.Println(user.UserId)
	if err := DB.Table("users").Where("user_id=?", user.UserId).Updates(map[string]interface{}{"user_name": user.UserName, "user_password": user.UserPassword, "user_picture": user.UserPicture, "motto": user.Motto}).Error; err != nil {
		return err
	}
	return nil
}

//删除我的发布中的书摘
func RemoveDigest(DigestId string) error {
	var summary DigestInfo
	if err1 := DB.Table("summaries").Where("id=?", DigestId).Delete(&summary).Error; err1 != nil {
		return err1
	}
	//var temp DigestAndClass
	//if err2 := DB.Table("summary_classes").Where("id=?", DigestId).Delete(&temp).Error; err2 != nil {
	//	return err2
	//}
	return nil
}

//删除书架上的书
func RemoveBook(BookId string) error {
	var UserBook UserAndBook
	if err := DB.Table("users_books").Where("book_id=?", BookId).Delete(&UserBook).Error; err != nil {
		return err
	}
	return nil
}

//查看书摘
func DigestPage(DigestId string) (DigestInfo, error) {
	var digest DigestInfo
	if err := DB.Table("summaries").Where("id=?", DigestId).Find(&digest).Error; err != nil {
		return DigestInfo{}, err
	}
	//fmt.Println(digest.DigestContent)
	return digest, nil
}

//查看图书
func BookPage(BookId string) (BooksInfo, error) {
	var book BooksInfo
	//fmt.Println(BookId)
	if err := DB.Table("books").Where("book_id=?", BookId).Find(&book).Error; err != nil {
		log.Println(err)
		return BooksInfo{}, err
	}
	return book, nil
}

func GetResult(content string) (BooksInfo, error) {
	var result BooksInfo
	//fmt.Println(content)
	if err := DB.Table("books").Where("book_name=?", content).Find(&result).Error; err != nil {
		return BooksInfo{}, err
	}
	return result, nil
}

//获取该书的全部评论
func GetAllDigest(BookId string) ([]DigestInfo, error) {
	var digest []DigestInfo
	if err := DB.Table("summaries").Where("book_id=?", BookId).Find(&digest).Error; err != nil {
		return nil, err
	}
	return digest, nil
}

//添加书本到我的书架里
func AddBookToShelf(BookId string, UserId string) error {
	var new UserAndBook
	new.UserId = UserId
	new.BookId = BookId
	if err := DB.Table("users_books").Create(&new).Error; err != nil {
		return err
	}
	return nil
}

//获取该书摘下的所有评论
func GetAllReview(DigestId string) ([]ReviewInfo, error) {
	var Reviews []ReviewInfo
	if err := DB.Table("reviews").Where("summary_id=?", DigestId).Find(&Reviews).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return Reviews, err
}

//获取评论界面用户的头像和昵称
func GetSomeThing(UserId string) (Userinfo, error) {
	var User Userinfo
	if err := DB.Table("users").Where("user_id=?", UserId).Find(&User).Error; err != nil {
		log.Println(err)
		return Userinfo{}, err
	}
	return User, nil
}

//创建一条评论
func CreatNewReview(UserId string, DigestId string, content string) error {
	temp := ReviewInfo{User_id: UserId, Summary_id: DigestId, Content: content}
	if err := DB.Table("reviews").Create(&temp).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}
