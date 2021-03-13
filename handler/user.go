package handler

import (
	"2021-Library-backend/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary "注册"
// @Description "注册一个新用户"
// @tags user
// @Accept json
// @Produce json
// @Success 200 "注册成功"
// @Failure 400 "输入有误，格式错误"
// @Router /user [post]
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
	if _, a := model.IfExistUserName(user.UserName); a != 1 {
		c.JSON(200, gin.H{
			"message": "该用户名已经存在请更换其他用户名"})
		return
	}
	user_id := model.Register(user.UserName, user.UserPassword)
	fmt.Println(user.UserName)
	c.JSON(200, gin.H{
		"user_id": user_id,
	})
}

// @Summary "登入"
// @Description "验证用户信息实现登入"
// @Tags login
// @Accept json
// @Producer json
// @Success 200 {object} model.Token "登入成功"
// @Failure 400 "输入格式错误"
// @Failure 404 "用户不存在"
// @Failure 401 "密码错误"
// @Router /login [post]
func Login(c *gin.Context) {
	var user model.Userinfo
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入格式错误"})
		return
	}
	fmt.Println(user.UserName, user.UserPassword)
	// 验证用户名是否存在
	if model.IfExistUser(user.UserName) != false {
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
		token := model.CreateToken(user.UserName, user.UserPassword, user.UserId)
		fmt.Println(token)
		return
	}
}
