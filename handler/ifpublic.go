package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func IfPublic(c *gin.Context) {
	var summary model.Summary
	summary_id := c.Query("summary_id")

	model.DB.First(&summary, summary_id)

	model.DB.Model(&summary).Update("public", !(summary.Public))
	c.JSON(200, gin.H{
		"message": "修改成功",
		"data":    summary.Public,
	})
}
