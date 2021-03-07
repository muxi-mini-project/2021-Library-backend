package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary "是否公开"
// @Description "通过按钮来确定是否公开这个书摘"
// @Tags digest
// @Accept json
// @Produce json
// @Success 200 "修改成功"
// @Failure 401 "修改失败"
// @Router /digest/mysummary/:user_id [put]
func IfPublic(c *gin.Context) {

	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	var summary model.Summary
	summary_id := c.Query("summary_id")

	model.DB.First(&summary, summary_id)

	model.DB.Model(&summary).Update("public", !(summary.Public))
	c.JSON(200, gin.H{
		"message": "修改成功",
		"data":    summary.Public,
	})
}
