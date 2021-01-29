package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户的信息"})
		return
	}
	MyName := model.GetUserName(id)
	MyPicture := model.GetUserPicture(id)
	MyMotto := model.GetUserMotto(id)
	userinfo := model.Userinfo{UserName: MyName, UserPicture: MyPicture, Motto: MyMotto}
	c.JSON(200, userinfo)
}

func Shelf(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户的信息"})
		return
	}
	Mybooks := model.GetMyBooksId(id)
	BooksThings, err1 := model.GetMyBooksinfo(Mybooks)
	if err1 != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, BooksThings)
}

func MyDigest(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该信息"})
		return
	}
	MyDigestInfo, err := model.GetMyDigest(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, MyDigestInfo)
}

//删除我的发布中的书摘
func DeleteDigest(c *gin.Context) {
	id := c.Param("digest_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}
	err1 := model.RemoveDigest(id)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

func Deletebooks(c *gin.Context) {
	id := c.Param("book_id")
	err := model.RemoveBook(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "删除失败"})
		return
	}
	c.JSON(200, gin.H{"message": "删除成功"})
}

//从我的页面查看我的书摘
func DigestId(c *gin.Context) {
	id := c.Param("digest_id")
	digest, err := model.DigestPage(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "查找失败"})
		return
	}
	c.JSON(200, digest)
}

//从我的页面查看我收藏的图书
func BookId(c *gin.Context) {
	id := c.Param("book_id")
	book, err := model.BookPage(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "查找失败"})
		return
	}
	c.JSON(200, book)
}
