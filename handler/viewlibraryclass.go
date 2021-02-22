package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func ViewLibraryClass(c *gin.Context) {
	var books []model.Book
	class_id := c.Param("class_id")

	model.DB.Where("class_id = ?", class_id).Find(&books)

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    books,
	})

}
