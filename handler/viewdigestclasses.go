package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary "书摘分类"
// @Description "书摘的各种类别"
// @Tags digest
// @Accept json
// @Produce json
// @Success 200 "获取成功"
// @Router /digest/mysummary/:user_id/classes [get]
func ViewDigestClasses(c *gin.Context) {

	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	var summaryClasses []model.SummaryClass
	user_id := c.Param("user_id")

	model.DB.Where("user_id = ?", user_id).Find(&summaryClasses)

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    summaryClasses,
	})

}
