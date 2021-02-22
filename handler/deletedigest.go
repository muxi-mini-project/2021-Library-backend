package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func DeleteDigest(c *gin.Context) {
	var summary model.Summary
	summary_id := c.Query("summary_id")

	model.DB.Delete(&summary, summary_id)

	c.JSON(200, gin.H{
		"message": "成功删除",
	})

}
