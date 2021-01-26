package handler

import (
	"github.com/gin-gonic/gin"
	"2021-Library-backend/model"
)

func User(c *gin.Context) {
	var user model.Userinfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400,gin.H{
			"message": "输入有误，格式错误"
		})
		return
	}
	user_id := model.Register(user.UserName, user.Password)	
	c.JSON(200, gin.H{
		"user_id" : user_id,
	})
}

func Login(c *gin.Context) {
	var usre model.userinfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400,gin.H{"message": "输入格式错误"})
		return
	}
	// 验证用户名是否存在
	if !model.IfExistUser(user.Id) {
		c.JSON(404, gin.H{"message": "用户不存在"})
		return
	}

	if !model.TestPassword(user.Id, user.Password) {
		c.JSON(401, gin.H{"message": "密码错误"})
		return
	}
}