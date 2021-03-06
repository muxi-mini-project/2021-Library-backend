package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary "书城"
// @Description "显示书城主界面"
// @Tags library
// @Accept json
// @Produce json
// @Success 200 "获取成功"
// @Router /library [get]
func ViewLibrary(c *gin.Context) {

	Spider()

	var books []model.Book
	var booksShow []model.Book
	num := 20 //推荐出的书目数

	model.DB.Find(&books)

	for i := len(books) - 1; i >= 0; i-- {
		if num > 0 {
			num--
		} else {
			break
		}
		booksShow = append(booksShow, books[i])
	}

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    booksShow,
	})

}
