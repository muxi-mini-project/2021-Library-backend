package handler

import(
	"study/model"
	"github.com/gin-gonic/gin"
)

func ViewDigest(c *gin.Context){
	var summaries []model.Summary
	user_id := c.Param("user_id")

	model.DB.Where("user_id = ?", string(user_id)).Find(&summaries)

	if len(summaries) <= 0 {
		c.JSON(200, gin.H{
			"message": "没有数据",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data": summaries,
	})
}
