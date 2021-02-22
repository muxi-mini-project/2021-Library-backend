package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func EditDigest(c *gin.Context) {
	var summary model.Summary
	var summaryInfo model.SummaryInfo
	var book model.Book
	summary_id := c.Param("summary_id")

	model.DB.First(&summary, summary_id)

	if summary.Id == 0 {
		c.JSON(404, gin.H{
			"message": "编辑失败，无数据",
		})
		return
	}

	err := c.BindJSON(&summaryInfo)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "编辑错误",
		})
		return
	}

	model.DB.Model(&summary).Update(map[string]interface{}{
		"title":               summaryInfo.Title,
		"chapter":             summaryInfo.Chapter,
		"summary_information": summaryInfo.Summary_information,
		"thought":             summaryInfo.Thought,
		"public":              summaryInfo.Public,
	})

	model.DB.Where("book_name = ?", summary.Title).First(&book)
	/*
		if book.Book_id == 0 {
			book.Book_id = 1
		}
	*/

	model.DB.Model(&summary).Update("book_id", book.Book_id)

	c.JSON(200, gin.H{
		"message": "编辑成功",
		"data":    summary,
	})

}
