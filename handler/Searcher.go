package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func Searcher(c *gin.Context) {
	var temp model.Search
	if err := c.BindJSON(&temp); err != nil {
		c.JSON(401, gin.H{"message": "请重试"})
		return
	}
	result, err2 := model.GetResult(temp.Content)
	if err2 != nil {
		c.JSON(404, gin.H{"message": "搜索不到"})
		return
	}
	c.JSON(200, result)
}

func BookId2(c *gin.Context) {
	id := c.Param("book_id")
	book, err := model.BookPage(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "查找失败"})
		return
	}
	c.JSON(200, book)
}

func Digest(c *gin.Context) {
	id := c.Param("book_id")
	DigestInfomation, err := model.GetAllDigest(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
		return
	}
	c.JSON(200, DigestInfomation)
}

func AddBook(c *gin.Context) {
	id := c.Param("book_id")
	token := c.Request.Header.Get("token")
	result, err1 := model.VerifyToken(token)
	if err1 != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	userid := result.UserId
	err := model.AddBookToShelf(id, userid)
	if err != nil {
		c.JSON(400, gin.H{"message": "添加失败"})
		return
	}
	c.JSON(200, gin.H{"message": "添加成功"})
	return
}
