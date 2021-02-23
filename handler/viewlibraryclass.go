package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary "书籍分类"
// @Description "书本的各种类别"
// @Tags library
// @Accept json
// @Produce json
// @Success 200 "获取成功"
// @Router /library/classes [get]
func ViewLibraryClass(c *gin.Context) {
	var books []model.Book
	class_id := c.Param("class_id")

	model.DB.Where("class_id = ?", class_id).Find(&books)

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    books,
	})

}
