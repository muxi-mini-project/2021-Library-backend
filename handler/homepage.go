package handler

import (
	"2021-Library-backend/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary "我的主页"
// @Description "通过我的id进入我的主页"
// @Tags my
// @Id get-string-by-int
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} model.Userinfo"{"msg":"success"}
// @Failure 404 "找不到该用户的信息"
// @Router /homepage/:user_id [get]
func HomePage(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户的信息"})
		return
	} //id := c.Param("user_id")
	//token := c.Request.Header.Get("token")
	//_, err := model.VerifyToken(token)
	//if err != nil {
	//	c.JSON(404, gin.H{"message": "找不到该用户的信息"})
	//	return
	//}
	fmt.Println(id)
	MyName := model.GetUserName(id)
	MyPicture := model.GetUserPicture(id)
	MyMotto := model.GetUserMotto(id)
	userinfo := model.Userinfo{UserName: MyName, UserPicture: MyPicture, Motto: MyMotto}
	c.JSON(200, userinfo)
}

// @Summary "我的书架"
// @Description "通过用户的id获得该用户的书架上书本的信息"
// @Tags my
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Success 200 {object} model.BooksInfo"{"msg":"success"}"
// @Failure 404 "获取失败"
// @Router /homepage/:user_id/shelf [get]
func Shelf(c *gin.Context) {
	//id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户的信息"})
		return
	}
	fmt.Println(id)
	Mybooks := model.GetMyBooksId(id)
	fmt.Println(Mybooks)
	BooksThings, err1 := model.GetMyBooksinfo(Mybooks)
	if err1 != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, BooksThings)
}

// @Summary "我的发布"
// @Description "显示我的发布中的书摘的有关信息"
// @Tags my
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Success 200 {object} model.DigestInfo "{"msg":"success"}"
// @Failure 404 "获取失败"
// @Router /homepage/:user_id/mydigest [get]
func MyDigest(c *gin.Context) {
	//id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该信息"})
		return
	}
	MyDigestInfo, err := model.GetMyDigest(id)
	//for _, x := range MyDigestInfo {
	//fmt.Println(x.UserId)
	//}
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, MyDigestInfo)
}

// @Summary "删除我的发布"
// @Description "删除我的发布中的书摘"
// @Tags my
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Success 200 "删除成功"
// @Failure 404 "认证失败"
// @Failure 401 "删除失败"
// @Router /homepage/:user_id/mydigest/digest_id [put]
func DeleteDigest(c *gin.Context) {
	Digestid := c.Param("digest_id")
	//fmt.Println(Digestid)
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}
	err1 := model.RemoveDigest(Digestid)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

// @Summary "删除我的书架上的书"
// @Description "删除数据库中这本书与该用户的关系"
// @Tags my
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Success 200 "删除成功"
// @Failure 401 "删除失败"
// @Router /homepage/:user_id/shelf/books_id [put]
func Deletebooks(c *gin.Context) {
	id := c.Param("books_id")
	err := model.RemoveBook(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

// @Summary "查看我的发布中书摘的具体信息"
// @Description "获取我的书=书摘的具体信息"
// @Tags my
// @Accept json
// @Produce json
// @Success 200 {object} model.DigestInfo
// @Failure 401 "获取失败"
// @Router /homepage/:user_id/mydigest/digest_id [get]
func DigestId(c *gin.Context) {
	id := c.Param("digest_id")
	digest, err := model.DigestPage(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "查找失败"})
		return
	}
	c.JSON(200, digest)
}

// @Summary "我收藏的书"
// @Description "查看我的书本的具体信息"
// @Tags my
// @Accept json
// @Producer json
// @Success 200 {object} model.BooksInfo
// @Failure 401 "查找失败"
// @Router /homepage/:user_id/shelf/book_id [get]
func BookId(c *gin.Context) {
	id := c.Param("books_id")
	book, err := model.BookPage(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "查找失败"})
		return
	}
	c.JSON(200, book)
}
