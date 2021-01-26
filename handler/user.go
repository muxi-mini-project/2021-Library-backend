package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/2021-Library-backend/model"
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