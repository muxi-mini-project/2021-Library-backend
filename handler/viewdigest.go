package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func ViewDigest(c *gin.Context) {
	var summaries []model.Summary
	var summaryRows []model.SummaryRow
	var summaryRow model.SummaryRow
	user_id := c.Param("user_id")

	model.DB.Where("user_id = ?", string(user_id)).Find(&summaries)

	if len(summaries) <= 0 {
		c.JSON(200, gin.H{
			"message": "没有数据",
		})
		return
	}

	for i := 0; i < len(summaries); i++ {
		summaryRow.Id = summaries[i].Id
		summaryRow.Title = summaries[i].Title
		summaryRow.Date = summaries[i].Date[0:10]
		summaryRow.Public = summaries[i].Public

		summaryRows = append(summaryRows, summaryRow)
	}

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    summaryRows,
	})
}
