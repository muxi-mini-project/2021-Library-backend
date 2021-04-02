package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary 删除书摘的分类
// @Description "删除我的书摘分类里的类别"
// @Tags digest
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param user_id path string true "user_id"
// @Success 200 "成功删除"
// @Router /digest/mysummary/{user_id}/clasees_edit [delete]
func DeleteDigestClass(c *gin.Context) {

	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	var summaryClass model.SummaryClass
	var summaries []model.Summary
	class_id := c.Query("class_id")

	model.DB.Where("class_id = ?", class_id).Delete(&summaryClass)
	model.DB.Where("class_id = ?", class_id).Delete(&summaries)

	c.JSON(200, gin.H{
		"message": "成功删除",
	})

}
