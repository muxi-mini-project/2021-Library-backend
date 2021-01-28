package handler

import (
	"2021-Library-backend/model"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户的信息"})
		return
	}
	MyName := model.GetUserName(id)
	MyPicture := model.GetUserPicture(id)
	MyMotto := model.GetUserMotto(id)
	userinfo := model.Userinfo{UserName: MyName, UserPicture: MyPicture, Motto: MyMotto}
	c.JSON(200, userinfo)
}

/* func Shelf(c *gin.Context) {
	id := c.Param("user_id")
	token := c.Request.Header.Get("token")
	_, err := model.VerifyToken(token)
	if err != nil {
		c.JSON(404, gin.H{"message": "找不到该用户的信息"})
		return
	}
	Mybooks := model.GetMyBooksId(id)
}
*/
