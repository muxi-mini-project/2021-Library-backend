package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary "获取我的书摘分类"
// @Descriptin "通过用户的id获得该用户的书摘分类"
// @Tags digest
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user_id path string true "user_id"
// @Success 200 {object} []model.SummaryClass "获取成功"
// @Router /digest/mysummary/{user_id}/classes_edit [get]
func EditDigestClasses(c *gin.Context) {

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
