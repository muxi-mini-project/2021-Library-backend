package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func DeleteDigestClass(c *gin.Context) {
	var summaryClass model.SummaryClass
	var summaries []model.Summary
	class_id := c.Query("class_id")

	model.DB.Where("class_id = ?", class_id).Delete(&summaryClass)
	model.DB.Where("class_id = ?", class_id).Delete(&summaries)

	c.JSON(200, gin.H{
		"message": "成功删除",
	})

}
