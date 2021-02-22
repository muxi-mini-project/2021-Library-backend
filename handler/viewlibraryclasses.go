package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func ViewLibraryClasses(c *gin.Context) {
	var bookClasses []model.BookClass
	var books []model.Book

	model.DB.Find(&bookClasses)
	model.DB.Find(&books)

	for i := 0; i < len(bookClasses); i++ {
		sum := 0
		for j := 0; j < len(books); j++ {
			if books[j].Class_id == bookClasses[i].Class_id {
				sum++
			}
		}
		bookClasses[i].Book_sum = int64(sum)
	}

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    bookClasses,
	})

}
