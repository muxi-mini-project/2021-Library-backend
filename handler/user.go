package handler

import (
	"2021-Library-backend/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	type users struct {
		UserName     string `json:"user_name"`
		UserPassword string `json:"user_password"`
	}
	var user users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}
	user_id := model.Register(user.UserName, user.UserPassword)
	fmt.Println(user.UserName)
	c.JSON(200, gin.H{
		"user_id": user_id,
	})
}

func Login(c *gin.Context) {
	var user model.Userinfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入格式错误"})
		return
	}
	fmt.Println(user.UserName, user.UserPassword)
	// 验证用户名是否存在
	if model.IfExistUser(user.UserName) == false {
		c.JSON(404, gin.H{"message": "用户不存在"})
		return
	}

	if model.TestPassword(user.UserName, user.UserPassword) == false {
		c.JSON(401, gin.H{"message": "密码错误"})
		return
	} else {
		user.UserId = model.GetId(user.UserName, user.UserPassword)
		c.JSON(200, gin.H{
			"message": "登录成功",
			"token":   model.CreateToken(user.UserName, user.UserPassword, user.UserId),
		})
		return
	}
}
