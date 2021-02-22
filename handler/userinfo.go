package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func Userinfo(c *gin.Context) {
	id := c.Param("user_id")
	Userinformation, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, Userinformation)
}

//修改密码的 hendel
func ChangeInformation(c *gin.Context) {
	var user model.Userinfo
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	if err1 := c.BindJSON(&user); err1 != nil {
		c.JSON(400, gin.H{"message": "输入格式有误"})
		return
	}
	user.UserId = id
	if err2 := model.ChangeUserInfo(user); err2 != nil {
		c.JSON(400, gin.H{"message": "修改失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}
