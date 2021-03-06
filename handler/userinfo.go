package handler

import (
	"2021-Library-backend/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

// @Summary "用户界面"
// @Description "获取用户的基本信息"
// @Tags my
// @Accept json
// @Produce json
// @Success 200 {object} model.Userinfo
// @Failure 404 "获取失败"
// Router /homepage/:user_id/info [get]
func Userinfo(c *gin.Context) {
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	Userinformation, err := model.GetUserInfo(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "获取失败"})
	}
	c.JSON(200, Userinformation)
}

// @Summary "修改用户的信息"
// @Description "修改用户的基本信息"
// @Tags my
// @Accept json
// @Produce json
// @Param toekn header string true "token"
// @Success 200 "修改成功"
// @Failure 401 "验证失败"
// @Failure 400 "修改失败"
// @Router /homepage/:user_id/info [put]
func ChangeInformation(c *gin.Context) {
	var user model.Userinfo
	//id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	id, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(401, gin.H{"message": "验证失败"})
	}
	if err1 := c.BindJSON(&user); err1 != nil {
		c.JSON(400, gin.H{"message": "输入格式有误"})
		return
	}
	user.UserId = id
	fmt.Println(user.UserId)
	if err2 := model.ChangeUserInfo(user); err2 != nil {
		c.JSON(400, gin.H{"message": "修改失败"})
		return
	}
	c.JSON(200, gin.H{"message": "修改成功"})
}