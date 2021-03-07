package handler

import (
	"strconv"

	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

// @Summary 添加书摘的分类
// @Description "添加属于自己的书摘分类"
// @Tags digest
// @Accept json
// @Produce json
// @Success 200 "添加成功"
// @Failure 400 "编辑错误"
// @Router /digest/mysummary/:user_id/classes_add [post]

func AddDigestClass(c *gin.Context) {

	token := c.Request.Header.Get("token")
	_, err0 := model.VerifyToken(token)
	if err0 != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	var summaryClass model.SummaryClass
	var summaryClassName model.SummaryClassName

	user_id := c.Param("user_id")
	summaryClass.User_id, _ = strconv.ParseInt(user_id, 10, 64)

	err := c.BindJSON(&summaryClassName)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "编辑错误",
		})
		return
	}

	summaryClass.Class_name = summaryClassName.Name

	model.DB.Create(&summaryClass)

	c.JSON(200, gin.H{
		"message": "添加成功",
		"data":    summaryClass,
	})

}
