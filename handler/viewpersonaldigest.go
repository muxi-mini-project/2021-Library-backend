﻿package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func ViewPersonalDigest(c *gin.Context) {

	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "认证失败"})
		return
	}

	var summary model.Summary
	summary_id := c.Param("summary_id")

	model.DB.First(&summary, summary_id)

	if summary.Id == 0 {
		c.JSON(404, gin.H{
			"message": "获取失败",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "获取成功",
		"data":    summary,
	})

}
