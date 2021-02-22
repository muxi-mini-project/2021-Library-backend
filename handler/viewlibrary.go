package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func ViewLibrary(c *gin.Context) {
	var books []model.Book
	var booksShow []model.Book
	num := 15 //推荐出的书目数

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
