package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary "搜索书本"
// @Description "在书城中搜索书本"
// @Tags library
// @Accept json
// @Produce json
// @Success 200 {object} model.BooksInfo "搜索成功"
// @Failure 401 "请重试"
// @Failure 404 "搜索不到"
// @Router /Library/seacher [post]
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

// @Summary "书籍信息"
// @Description "从书城进入书籍"
// @Tags library
// @Accept json
// @Produce json
// @Success 200 {object} model.BooksInfo
// @Failure 401 "查找失败"
// @Router /Library/:books_id [get]
func BookId2(c *gin.Context) {
	id := c.Param("books_id")
	book, err := model.BookPage(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "查找失败"})
		return
	}
	c.JSON(200, book)
}

// @Summary "书本的书摘信息"
// @Description "从书城进入书本获取的关于这个书的书摘信息"
// @Tags library
// @Accept json
// @Produce json
// @Success 200 {object} model.DigestInfo
// @Failure 404 "获取失败"
// @Router /Library/:books_id/digest [get]
func Digest(c *gin.Context) {
	id := c.Param("books_id")
	DigestInfomation, err := model.GetAllDigest(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
		return
	}
	c.JSON(200, DigestInfomation)
}

// @Summary "添加书本到我的书架"
// @Description "通过书本上面的添加按钮来实现添加"
// @Tags library
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Success 200 "添加成功"
// @Failure 401 "验证失败"
// @Failure 400 "添加失败"
// @Router /Library/addbook/:books_id [post]
func AddBook(c *gin.Context) {
	id := c.Param("books_id")
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
